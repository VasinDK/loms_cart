package server

import "net/http"

func (s *Server) CleaningCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getPathValueInt(w, r, "user_id")
	if err != nil {
		return
	}

	err = s.Service.ClearCart(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
