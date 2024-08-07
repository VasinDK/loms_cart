// Code generated by http://github.com/gojuno/minimock (v3.3.11). DO NOT EDIT.

package mock

//go:generate minimock -i route256/cart/internal/service/list/clear_cart.Repository -o repository_mock.go -n RepositoryMock -p mock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RepositoryMock implements clear_cart.Repository
type RepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcClearCart          func(ctx context.Context, i1 int64) (err error)
	inspectFuncClearCart   func(ctx context.Context, i1 int64)
	afterClearCartCounter  uint64
	beforeClearCartCounter uint64
	ClearCartMock          mRepositoryMockClearCart
}

// NewRepositoryMock returns a mock for clear_cart.Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ClearCartMock = mRepositoryMockClearCart{mock: m}
	m.ClearCartMock.callArgs = []*RepositoryMockClearCartParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRepositoryMockClearCart struct {
	optional           bool
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockClearCartExpectation
	expectations       []*RepositoryMockClearCartExpectation

	callArgs []*RepositoryMockClearCartParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// RepositoryMockClearCartExpectation specifies expectation struct of the Repository.ClearCart
type RepositoryMockClearCartExpectation struct {
	mock      *RepositoryMock
	params    *RepositoryMockClearCartParams
	paramPtrs *RepositoryMockClearCartParamPtrs
	results   *RepositoryMockClearCartResults
	Counter   uint64
}

// RepositoryMockClearCartParams contains parameters of the Repository.ClearCart
type RepositoryMockClearCartParams struct {
	ctx context.Context
	i1  int64
}

// RepositoryMockClearCartParamPtrs contains pointers to parameters of the Repository.ClearCart
type RepositoryMockClearCartParamPtrs struct {
	ctx *context.Context
	i1  *int64
}

// RepositoryMockClearCartResults contains results of the Repository.ClearCart
type RepositoryMockClearCartResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option by default unless you really need it, as it helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmClearCart *mRepositoryMockClearCart) Optional() *mRepositoryMockClearCart {
	mmClearCart.optional = true
	return mmClearCart
}

// Expect sets up expected params for Repository.ClearCart
func (mmClearCart *mRepositoryMockClearCart) Expect(ctx context.Context, i1 int64) *mRepositoryMockClearCart {
	if mmClearCart.mock.funcClearCart != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Set")
	}

	if mmClearCart.defaultExpectation == nil {
		mmClearCart.defaultExpectation = &RepositoryMockClearCartExpectation{}
	}

	if mmClearCart.defaultExpectation.paramPtrs != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by ExpectParams functions")
	}

	mmClearCart.defaultExpectation.params = &RepositoryMockClearCartParams{ctx, i1}
	for _, e := range mmClearCart.expectations {
		if minimock.Equal(e.params, mmClearCart.defaultExpectation.params) {
			mmClearCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmClearCart.defaultExpectation.params)
		}
	}

	return mmClearCart
}

// ExpectCtxParam1 sets up expected param ctx for Repository.ClearCart
func (mmClearCart *mRepositoryMockClearCart) ExpectCtxParam1(ctx context.Context) *mRepositoryMockClearCart {
	if mmClearCart.mock.funcClearCart != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Set")
	}

	if mmClearCart.defaultExpectation == nil {
		mmClearCart.defaultExpectation = &RepositoryMockClearCartExpectation{}
	}

	if mmClearCart.defaultExpectation.params != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Expect")
	}

	if mmClearCart.defaultExpectation.paramPtrs == nil {
		mmClearCart.defaultExpectation.paramPtrs = &RepositoryMockClearCartParamPtrs{}
	}
	mmClearCart.defaultExpectation.paramPtrs.ctx = &ctx

	return mmClearCart
}

// ExpectI1Param2 sets up expected param i1 for Repository.ClearCart
func (mmClearCart *mRepositoryMockClearCart) ExpectI1Param2(i1 int64) *mRepositoryMockClearCart {
	if mmClearCart.mock.funcClearCart != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Set")
	}

	if mmClearCart.defaultExpectation == nil {
		mmClearCart.defaultExpectation = &RepositoryMockClearCartExpectation{}
	}

	if mmClearCart.defaultExpectation.params != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Expect")
	}

	if mmClearCart.defaultExpectation.paramPtrs == nil {
		mmClearCart.defaultExpectation.paramPtrs = &RepositoryMockClearCartParamPtrs{}
	}
	mmClearCart.defaultExpectation.paramPtrs.i1 = &i1

	return mmClearCart
}

// Inspect accepts an inspector function that has same arguments as the Repository.ClearCart
func (mmClearCart *mRepositoryMockClearCart) Inspect(f func(ctx context.Context, i1 int64)) *mRepositoryMockClearCart {
	if mmClearCart.mock.inspectFuncClearCart != nil {
		mmClearCart.mock.t.Fatalf("Inspect function is already set for RepositoryMock.ClearCart")
	}

	mmClearCart.mock.inspectFuncClearCart = f

	return mmClearCart
}

// Return sets up results that will be returned by Repository.ClearCart
func (mmClearCart *mRepositoryMockClearCart) Return(err error) *RepositoryMock {
	if mmClearCart.mock.funcClearCart != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Set")
	}

	if mmClearCart.defaultExpectation == nil {
		mmClearCart.defaultExpectation = &RepositoryMockClearCartExpectation{mock: mmClearCart.mock}
	}
	mmClearCart.defaultExpectation.results = &RepositoryMockClearCartResults{err}
	return mmClearCart.mock
}

// Set uses given function f to mock the Repository.ClearCart method
func (mmClearCart *mRepositoryMockClearCart) Set(f func(ctx context.Context, i1 int64) (err error)) *RepositoryMock {
	if mmClearCart.defaultExpectation != nil {
		mmClearCart.mock.t.Fatalf("Default expectation is already set for the Repository.ClearCart method")
	}

	if len(mmClearCart.expectations) > 0 {
		mmClearCart.mock.t.Fatalf("Some expectations are already set for the Repository.ClearCart method")
	}

	mmClearCart.mock.funcClearCart = f
	return mmClearCart.mock
}

// When sets expectation for the Repository.ClearCart which will trigger the result defined by the following
// Then helper
func (mmClearCart *mRepositoryMockClearCart) When(ctx context.Context, i1 int64) *RepositoryMockClearCartExpectation {
	if mmClearCart.mock.funcClearCart != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Set")
	}

	expectation := &RepositoryMockClearCartExpectation{
		mock:   mmClearCart.mock,
		params: &RepositoryMockClearCartParams{ctx, i1},
	}
	mmClearCart.expectations = append(mmClearCart.expectations, expectation)
	return expectation
}

// Then sets up Repository.ClearCart return parameters for the expectation previously defined by the When method
func (e *RepositoryMockClearCartExpectation) Then(err error) *RepositoryMock {
	e.results = &RepositoryMockClearCartResults{err}
	return e.mock
}

// Times sets number of times Repository.ClearCart should be invoked
func (mmClearCart *mRepositoryMockClearCart) Times(n uint64) *mRepositoryMockClearCart {
	if n == 0 {
		mmClearCart.mock.t.Fatalf("Times of RepositoryMock.ClearCart mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmClearCart.expectedInvocations, n)
	return mmClearCart
}

func (mmClearCart *mRepositoryMockClearCart) invocationsDone() bool {
	if len(mmClearCart.expectations) == 0 && mmClearCart.defaultExpectation == nil && mmClearCart.mock.funcClearCart == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmClearCart.mock.afterClearCartCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmClearCart.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// ClearCart implements clear_cart.Repository
func (mmClearCart *RepositoryMock) ClearCart(ctx context.Context, i1 int64) (err error) {
	mm_atomic.AddUint64(&mmClearCart.beforeClearCartCounter, 1)
	defer mm_atomic.AddUint64(&mmClearCart.afterClearCartCounter, 1)

	if mmClearCart.inspectFuncClearCart != nil {
		mmClearCart.inspectFuncClearCart(ctx, i1)
	}

	mm_params := RepositoryMockClearCartParams{ctx, i1}

	// Record call args
	mmClearCart.ClearCartMock.mutex.Lock()
	mmClearCart.ClearCartMock.callArgs = append(mmClearCart.ClearCartMock.callArgs, &mm_params)
	mmClearCart.ClearCartMock.mutex.Unlock()

	for _, e := range mmClearCart.ClearCartMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmClearCart.ClearCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClearCart.ClearCartMock.defaultExpectation.Counter, 1)
		mm_want := mmClearCart.ClearCartMock.defaultExpectation.params
		mm_want_ptrs := mmClearCart.ClearCartMock.defaultExpectation.paramPtrs

		mm_got := RepositoryMockClearCartParams{ctx, i1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmClearCart.t.Errorf("RepositoryMock.ClearCart got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.i1 != nil && !minimock.Equal(*mm_want_ptrs.i1, mm_got.i1) {
				mmClearCart.t.Errorf("RepositoryMock.ClearCart got unexpected parameter i1, want: %#v, got: %#v%s\n", *mm_want_ptrs.i1, mm_got.i1, minimock.Diff(*mm_want_ptrs.i1, mm_got.i1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmClearCart.t.Errorf("RepositoryMock.ClearCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmClearCart.ClearCartMock.defaultExpectation.results
		if mm_results == nil {
			mmClearCart.t.Fatal("No results are set for the RepositoryMock.ClearCart")
		}
		return (*mm_results).err
	}
	if mmClearCart.funcClearCart != nil {
		return mmClearCart.funcClearCart(ctx, i1)
	}
	mmClearCart.t.Fatalf("Unexpected call to RepositoryMock.ClearCart. %v %v", ctx, i1)
	return
}

// ClearCartAfterCounter returns a count of finished RepositoryMock.ClearCart invocations
func (mmClearCart *RepositoryMock) ClearCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClearCart.afterClearCartCounter)
}

// ClearCartBeforeCounter returns a count of RepositoryMock.ClearCart invocations
func (mmClearCart *RepositoryMock) ClearCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClearCart.beforeClearCartCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.ClearCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmClearCart *mRepositoryMockClearCart) Calls() []*RepositoryMockClearCartParams {
	mmClearCart.mutex.RLock()

	argCopy := make([]*RepositoryMockClearCartParams, len(mmClearCart.callArgs))
	copy(argCopy, mmClearCart.callArgs)

	mmClearCart.mutex.RUnlock()

	return argCopy
}

// MinimockClearCartDone returns true if the count of the ClearCart invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockClearCartDone() bool {
	if m.ClearCartMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ClearCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ClearCartMock.invocationsDone()
}

// MinimockClearCartInspect logs each unmet expectation
func (m *RepositoryMock) MinimockClearCartInspect() {
	for _, e := range m.ClearCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.ClearCart with params: %#v", *e.params)
		}
	}

	afterClearCartCounter := mm_atomic.LoadUint64(&m.afterClearCartCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ClearCartMock.defaultExpectation != nil && afterClearCartCounter < 1 {
		if m.ClearCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.ClearCart")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.ClearCart with params: %#v", *m.ClearCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearCart != nil && afterClearCartCounter < 1 {
		m.t.Error("Expected call to RepositoryMock.ClearCart")
	}

	if !m.ClearCartMock.invocationsDone() && afterClearCartCounter > 0 {
		m.t.Errorf("Expected %d calls to RepositoryMock.ClearCart but found %d calls",
			mm_atomic.LoadUint64(&m.ClearCartMock.expectedInvocations), afterClearCartCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockClearCartInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockClearCartDone()
}
