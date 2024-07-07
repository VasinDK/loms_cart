package http_handlers

import (
	"context"
	"net/http"
	"route256/cart/internal/pkg/logger"
	"strconv"
)

// getPathValueInt - извлекает параметры из строки, меняя тип, проверяя ошибки.
// В случае ошибки отвечает 500
func getPathValueInt(w http.ResponseWriter, r *http.Request, param string) (int64, error) {
	paramStr := r.PathValue(param)
	paramNum, err := strconv.Atoi(paramStr)
	ctx := context.Background()

	if err != nil {
		logger.Errorw(ctx, "getPathValueInt", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return 0, err
	}

	return int64(paramNum), nil
}
