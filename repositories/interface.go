package repositories

import "library-system/models"

type BookRepository interface {
	Create(book *models.Book) error
}

type BorrowerRepository interface {
	Create(borrower *models.Borrower) error
}

type HistoryRepository interface {
	Create(history *models.HistoryBorrower) error
	Update(history *models.HistoryBorrower) error
	GetHistory(offset, pageSize int) ([]models.HistoryBorrower, error)
	GetHistoryById(id uint) (*models.HistoryBorrower, error)
	FindOverdueBooks(offset, pageSize int) ([]models.HistoryBorrower, error)
	FindMostBorrowedBooks(offset, pageSize int) ([]models.Book, error)
}
