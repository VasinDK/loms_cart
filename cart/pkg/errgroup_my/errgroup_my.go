package errgroup_my

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Elem struct{}

// CountInPeriod - Учет количества горутин работающих в данный период времени
type CountInPeriod struct {
	max             int   // Максимальное количество горутин запускаемых в период
	current         int   // Текущее количество горутин в текущем периоде
	deadline        int64 // Окончание текущего периода в таймстемп
	secondsInPeriod int   // Количество секунд в периоде
}

type Group struct {
	cancel        func(error)
	wg            sync.WaitGroup
	err           error
	errOnce       sync.Once
	count         chan Elem // количество горутин работающих в данный момент
	countInPeriod CountInPeriod
}

// WithContext - создает экземпляр и ctx с Cause
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancelCause(ctx)
	return &Group{cancel: cancel}, ctx
}

// Go - Если лимит не задан то сразу добавляет горутины. В случае если задан лимит на горутины
// добавляет или ждет освобождения и потом добавляет новые горутины.
//
// В случае возникновения ошибки в обработчике сохраняет ошибку, вызывает
// cancel и отменяет ctx всех горутин
func (g *Group) Go(f func() error) {
	if g.count != nil {
		g.count <- Elem{}
	}

	if g.countInPeriod.max > 0 {
		ch1 := g.addItemInPeriod()
		<-ch1
	}

	g.wg.Add(1)
	go func() {
		defer g.done()
		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel(g.err)
				}
			})
		}
	}()
}

// Done - уменьшает счетчик ожидания горутин и занятость
func (g *Group) done() {
	if g.count != nil {
		<-g.count
	}
	g.wg.Done()
}

// Wait - ждет окончания работы всех горутин. В случае ошибок вызывает отмену
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel(g.err)
	}
	return g.err
}

// TryGo - аналогичен Go но не ждет освобождения, а проверяет и если есть место то запускает,
// если нет то отвечает false
func (g *Group) TryGo(f func() error) bool {
	if g.count != nil {
		select {
		case g.count <- Elem{}:
		default:
			return false
		}
	}

	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel(g.err)
				}
			})
		}
	}()
	return true
}

// SetLimit - устанавливает лимит количества одновременно работающих горутин
func (g *Group) SetLimit(n int) {
	if n < 0 {
		g.count = nil
		return
	}

	if len(g.count) != 0 {
		panic(fmt.Errorf("errgroup-my: Нельзя изменить лимит. %v горутин еще активны", len(g.count)))
	}

	g.count = make(chan Elem, n)
}

// SetLimitPeriod - устанавливает количество горутин разрешенных в период и сам период
func (g *Group) SetLimitPeriod(count int, seconds int) {
	if count < 0 {
		g.countInPeriod = CountInPeriod{}
		return
	}

	g.countInPeriod = CountInPeriod{
		max:             count,
		secondsInPeriod: seconds,
	}
}

// addItemInPeriod - Добавляет горутину в работу в текущий период
func (g *Group) addItemInPeriod() <-chan struct{} {
	ch1 := make(chan struct{})

	go func() {
		for {
			currentTime := time.Now().Unix()

			if g.countInPeriod.deadline < currentTime {
				g.countInPeriod.deadline = currentTime + int64(g.countInPeriod.secondsInPeriod)
				g.countInPeriod.current = 1
				ch1 <- struct{}{}
				break
			}

			if g.countInPeriod.current < g.countInPeriod.max {
				g.countInPeriod.current++
				ch1 <- struct{}{}
				break
			}
		}
		close(ch1)
	}()

	return ch1
}
