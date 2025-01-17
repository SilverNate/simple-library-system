package handlers

import (
	"net/http"
)

type HistoryHandler interface {
	CreateHistory(w http.ResponseWriter, r *http.Request)
	UpdateReturnDate(w http.ResponseWriter, r *http.Request)
	GetHistory(w http.ResponseWriter, r *http.Request)
	GetOverdueBooks(w http.ResponseWriter, r *http.Request)
	FindMostBorrowedBooks(w http.ResponseWriter, r *http.Request)
}

type BookHandler interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
}

type BorrowerHandler interface {
	CreateBorrower(w http.ResponseWriter, r *http.Request)
}
