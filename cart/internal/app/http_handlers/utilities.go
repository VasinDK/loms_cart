package http_handlers

import (
	"log/slog"
	"net/http"
	"strconv"
)

// Извлекает параметры из строки, меняя тип, проверяя ошибки.
// В случае ошибки отвечает 500
func getPathValueInt(w http.ResponseWriter, r *http.Request, param string) (int64, error) {
	paramStr := r.PathValue(param)
	paramNum, err := strconv.Atoi(paramStr)
	if err != nil {
		slog.Error("getPathValueInt", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return 0, err
	}

	return int64(paramNum), nil
}
