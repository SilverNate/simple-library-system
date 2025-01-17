package repositories

import (
	"gorm.io/gorm"
	"library-system/models"
	"log"
)

type historyRepository struct {
	db *gorm.DB
}

// creates a new instance of historyRepository
func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}

func (r *historyRepository) Create(history *models.HistoryBorrower) error {
	return r.db.Create(history).Error
}

func (r *historyRepository) GetHistory(offset, pageSize int) ([]models.HistoryBorrower, error) {
	var histories []models.HistoryBorrower
	err := r.db.Limit(pageSize).Offset(offset).Find(&histories).Error
	return histories, err
}

func (r *historyRepository) Update(history *models.HistoryBorrower) error {
	return r.db.Save(history).Error
}

func (r *historyRepository) GetHistoryById(id uint) (history *models.HistoryBorrower, err error) {
	err = r.db.First(&history, id).Error
	if err != nil {
		log.Printf("Error get history by id repository: %v", err)
		return history, err
	}
	return
}

func (r *historyRepository) FindOverdueBooks(offset, pageSize int) ([]models.HistoryBorrower, error) {
	var histories []models.HistoryBorrower
	result := r.db.Where("due_date < CURRENT_DATE AND return_date IS NULL").Limit(pageSize).Offset(offset).Find(&histories)
	return histories, result.Error
}

func (r *historyRepository) FindMostBorrowedBooks(offset, pageSize int) ([]models.Book, error) {
	var books []models.Book
	result := r.db.Raw(`
        SELECT b.*, COUNT(h.book_id) as borrow_count 
        FROM history_borrowers h
        JOIN books b ON b.id = h.book_id 
        GROUP BY b.id 
        ORDER BY borrow_count DESC LIMIT 5`).Limit(pageSize).Offset(offset).Scan(&books)
	return books, result.Error
}
