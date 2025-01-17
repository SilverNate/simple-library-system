package handlers

import (
	"encoding/json"
	"library-system/helper/utils"
	"library-system/models"
	"library-system/services"
	"net/http"
	"strconv"
)

type historyHandler struct {
	historyService services.HistoryService
}

// NewHistoryHandler creates a new instance of HistoryHandler interface
func NewHistoryHandler(service services.HistoryService) HistoryHandler {
	return &historyHandler{historyService: service}
}

func (h *historyHandler) CreateHistory(w http.ResponseWriter, r *http.Request) {
	var request models.RequestHistoryBorrower
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if err := h.historyService.AddHistoryRecord(request); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Error adding history record", err)
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, "history record added successfully", request)
}

func (h *historyHandler) UpdateReturnDate(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Id         int    `json:"id"`
		ReturnDate string `json:"return_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if err := h.historyService.UpdateReturnDate(uint(requestBody.Id), requestBody.ReturnDate); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Error update return date", err)
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "return date updated successfully", nil)
}

func (h *historyHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	histories, err := h.historyService.GetHistory(offset, pageSize)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Error get history records", err)
		return
	}

	utils.SendJSONResponseWithPagination(w, http.StatusOK, "get history records successfully", histories, page, pageSize)
}

func (h *historyHandler) GetOverdueBooks(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	overdueBooks, err := h.historyService.GetOverdueBooks(offset, pageSize)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Error get overdue books", err)
		return
	}

	// use simple pagination with pagesize and offset
	utils.SendJSONResponseWithPagination(w, http.StatusOK, "get overdue books successfully", overdueBooks, page, pageSize)
}

func (h *historyHandler) FindMostBorrowedBooks(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	overdueBooks, err := h.historyService.FindMostBorrowedBooks(offset, pageSize)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Error find most borrowed books", err)
		return
	}

	utils.SendJSONResponseWithPagination(w, http.StatusOK, "get most borrowed books successfully", overdueBooks, page, pageSize)
}
