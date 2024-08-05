package statuses

import (
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/status"
)

// GetCodePG - получает статус ошибки в виде кода
func GetCodePG(err error) string {
	var errLabel = "OK"

	if err == nil {
		return errLabel
	}

	if errPg, ok := err.(*pgconn.PgError); ok {
		return errPg.Code
	}

	return err.Error()
}

// GetCodeHTTP - получает статус ошибки в виде кода
func GetCodeHTTP(err error) string {
	var errLabel = "OK"

	if err == nil {
		return errLabel
	}

	if errHTTP, ok := err.(*url.Error); ok {
		return errHTTP.Err.Error()
	}

	return err.Error()
}

// GetStatusGRPC - получает статус ошибки
func GetStatusGRPC(err error) string {
	return status.Convert(err).String()
}

// GetStatusCodeGRPC - получает статус ошибки в виде кода
func GetStatusCodeGRPC(err error) string {
	return status.Convert(err).Code().String()
}

// GetStatusCodeRedis - получает статус ошибки redis в виде кода
func GetStatusCodeRedis(err error) string {
	if err == redis.Nil {
		return "no value"
	}

	if err != nil {
		if redisErr, ok := err.(redis.Error); ok {
			return redisErr.Error()
		}

		return "error"
	}

	return "ok"
}

func GetStatusCodeWebSocket(err error) string {
	if err == nil {
		return "ok"
	}

	if closeErr, ok := err.(*websocket.CloseError); ok {
		return closeErr.Error()
	}

	return "error"
}
