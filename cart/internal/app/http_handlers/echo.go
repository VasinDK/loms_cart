package http_handlers

import (
	"net/http"
	"route256/cart/internal/service/item/echo"
)

// Echo - Эхо реализованное на основе вебсокета
func (s *Server) Echo(ec *echo.Echo) http.Handler {
	return ec.Echo()
}
