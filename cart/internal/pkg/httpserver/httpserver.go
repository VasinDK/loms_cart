package httpserver

import (
	"log/slog"
	"net/http"
)

type Config interface {
	GetPort() string
}

func Run(mux http.Handler, config Config) {
	err := http.ListenAndServe(config.GetPort(), mux)
	if err != nil {
		slog.Error(err.Error())
	}
}
