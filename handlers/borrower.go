package handlers

import (
	"encoding/json"
	"library-system/helper/utils"
	"library-system/models"
	"library-system/services"
	"net/http"
)

type borrowerHandler struct {
	borrowerService services.BorrowerService
}

func NewBorrowerHandler(service services.BorrowerService) BorrowerHandler {
	return &borrowerHandler{borrowerService: service}
}

func (h *borrowerHandler) CreateBorrower(w http.ResponseWriter, r *http.Request) {
	var borrower models.Borrower
	json.NewDecoder(r.Body).Decode(&borrower)
	err := h.borrowerService.CreateBorrower(&borrower)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Error adding borrower", err)
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, "borrower added successfully", nil)
}
