package http_handlers

import (
	"net/http"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/service/item/echo"
)

// Echo - Эхо реализованное на основе вебсокета
// func (s *Server) Echo(ctx context.Context, ec *echo.Echo) http.Handler {
// return ec.Echo(ctx)
// }

func (s *Server) Echo(ec *echo.Echo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := ec.Echo(ctx)
		if err != nil {
			logger.Errorw(ctx, " ec.Echo", "err", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
