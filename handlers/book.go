package handlers

import (
	"encoding/json"
	"library-system/helper/utils"
	"library-system/models"
	"library-system/services"
	"net/http"
)

type bookHandler struct {
	bookService services.BookService
}

func NewBookHandler(service services.BookService) BookHandler {
	return &bookHandler{bookService: service}
}

func (h *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if err := h.bookService.CreateBook(&book); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Error adding book", err)
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, "book added successfully", nil)

}
