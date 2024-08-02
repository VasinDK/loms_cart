package echo

import (
	"context"
	"route256/cart/internal/pkg/logger"

	"golang.org/x/net/websocket"
)

type Echo struct{}

func New() *Echo {
	return &Echo{}
}

func (e *Echo) Echo() websocket.Handler {
	return websocket.Handler(run)
}

func run(conn *websocket.Conn) {
	ctx := context.Background()
	defer conn.Close()

	// Можно добавить ping. Прослушивать и закрывать.
	// Можно добавить закрытие по приходу syscall.SIGINT

	var msg string
	for {
		if err := websocket.Message.Receive(conn, &msg); err != nil {
			logger.Errorw(ctx, "websocket.Message.Receive", "err.Error", err.Error())
			break
		}

		if err := websocket.Message.Send(conn, msg); err != nil {
			logger.Errorw(ctx, "websocket.Message.Send", "err.Error", err.Error())
			break
		}
	}
}
