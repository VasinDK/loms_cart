package echo

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)

type Repo interface {
	GetByKeyMemDB(ctx context.Context, key string) (string, error)
	SetKeyMemDB(ctx context.Context, key string, value string) error
	ConnWebSocket(w http.ResponseWriter, req *http.Request, BufferSize int) (*websocket.Conn, error)
	ReadWebSocket()
	WriteWebSocket()
}

type Echo struct {
	Repo Repo
}

func New(repo Repo) *Echo {
	return &Echo{
		repo,
	}
}

func (e *Echo) Echo(ctx context.Context) error {
	/* return websocket.Handler(func(conn *websocket.Conn) {
		defer conn.Close()

		// Можно добавить ping. Прослушивать и закрывать.
		// Можно добавить закрытие по приходу syscall.SIGINT

		var msg string
		for {
			if err := websocket.Message.Receive(conn, &msg); err != nil {
				logger.Errorw(ctx, "websocket.Message.Receive", "err.Error", err.Error())
				break
			}

			keyMess := strconv.FormatInt(time.Now().Unix(), 10)

			err := e.Repo.SetKeyMemDB(ctx, keyMess, msg)
			if err != nil {
				logger.Errorw(ctx, "e.Repo.SetKeyMemDB", "err", err.Error())
			}

			value, err := e.Repo.GetByKeyMemDB(ctx, keyMess)
			if err != nil {
				logger.Errorw(ctx, "e.Repo.GetByKeyMemDB", "err", err.Error())
			}

			if err := websocket.Message.Send(conn, value); err != nil {
				logger.Errorw(ctx, "websocket.Message.Send", "err.Error", err.Error())
				break
			}
		}
	}) */

	// ConnWebSocket(w http.ResponseWriter, req *http.Request, BufferSize int) (*websocket.Conn, error)

}
