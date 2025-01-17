package handlers

import (
	"library-system/helper/utils"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	token, _ := utils.GenerateToken(email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token":"` + token + `"}`))
}
