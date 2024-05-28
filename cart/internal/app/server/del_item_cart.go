package server

import "net/http"

func (s *Server) DelItemCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getPathValueInt(w, r, "user_id")
	if err != nil {
		return
	}

	sku, err := getPathValueInt(w, r, "sku_id")
	if err != nil {
		return
	}

	err = s.Service.DeleteSKU(userId, sku)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
