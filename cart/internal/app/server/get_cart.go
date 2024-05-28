package server

import "net/http"

func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getPathValueInt(w, r, "user_id")
	if err != nil {
		return
	}

	_, err = s.Service.GetCart(userId)
	
}
