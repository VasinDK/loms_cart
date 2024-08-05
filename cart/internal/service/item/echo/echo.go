package echo

import (
	"context"
	"fmt"
	"net/http"
	"route256/cart/internal/pkg/logger"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type Repo interface {
	GetByKeyMemDB(ctx context.Context, key string) (string, error)
	SetKeyMemDB(ctx context.Context, key string, value string) error
	ConnWebSocket(ctx context.Context, w http.ResponseWriter, req *http.Request) (*websocket.Conn, error)
	ReadWebSocket(conn *websocket.Conn) (string, error)
	WriteWebSocket(conn *websocket.Conn, value string) error
}

type Echo struct {
	Repo Repo
}

// New - новый экземпляр сервиса Echo
func New(repo Repo) *Echo {
	return &Echo{
		repo,
	}
}

// Echo - принимает сообщение записывает в реди, возвращает обратно
func (e *Echo) Echo(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	conn, err := e.Repo.ConnWebSocket(ctx, w, req)
	if err != nil {
		return fmt.Errorf("e.Repo.ConnWebSocket %w", err)
	}
	defer conn.Close()

	// Можно добавить ping. Прослушивать и закрывать.
	// Можно добавить закрытие по приходу syscall.SIGINT

	for {
		msg, err := e.Repo.ReadWebSocket(conn)
		if err != nil {
			logger.Errorw(ctx, "e.Repo.ReadWebSocket.", "err", err.Error())
			break
		}

		keyMess := strconv.FormatInt(time.Now().Unix(), 10)

		err = e.Repo.SetKeyMemDB(ctx, keyMess, msg)
		if err != nil {
			logger.Errorw(ctx, "e.Repo.SetKeyMemDB", "err", err.Error())
		}

		value, err := e.Repo.GetByKeyMemDB(ctx, keyMess)
		if err != nil {
			logger.Errorw(ctx, "e.Repo.GetByKeyMemDB", "err", err.Error())
		}

		err = e.Repo.WriteWebSocket(conn, value)
		if err != nil {
			logger.Errorw(ctx, "websocket.Message.Send", "err.Error", err.Error())
			break
		}
	}

	return nil
}
