package server

import (
	"net/http"
	"strconv"
)

func getPathValueInt(w http.ResponseWriter, r *http.Request, param string) (int64, error) {
	paramStr := r.PathValue(param)
	paramNum, err := strconv.Atoi(paramStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return 0, err
	}

	return int64(paramNum), nil
}
