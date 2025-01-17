package repositories

import (
	"gorm.io/gorm"
	"library-system/models"
)

type borrowerRepository struct {
	db *gorm.DB
}

func NewBorrowerRepository(db *gorm.DB) BorrowerRepository {
	return &borrowerRepository{db: db}
}

func (r borrowerRepository) Create(borrower *models.Borrower) error {
	return r.db.Create(borrower).Error
}
