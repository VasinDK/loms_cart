package statuses

import (
	"net/url"

	"github.com/jackc/pgx/v5/pgconn"
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
