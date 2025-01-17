package services

import "library-system/models"

type BookService interface {
	CreateBook(book *models.Book) error
}

type BorrowerService interface {
	CreateBorrower(borrower *models.Borrower) error
}

type HistoryService interface {
	AddHistoryRecord(history models.RequestHistoryBorrower) error
	UpdateReturnDate(id uint, returnDate string) error
	GetHistory(offset, pageSize int) ([]models.HistoryBorrower, error)
	GetOverdueBooks(offset, pageSize int) ([]models.HistoryBorrower, error)
	FindMostBorrowedBooks(offset, pageSize int) ([]models.Book, error)
}
