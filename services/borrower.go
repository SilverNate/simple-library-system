package services

import (
	"library-system/models"
	"library-system/repositories"
)

type borrowerService struct {
	borrowerRepo repositories.BorrowerRepository
}

func NewBorrowerService(repo repositories.BorrowerRepository) BorrowerService {
	return &borrowerService{borrowerRepo: repo}
}

func (s borrowerService) CreateBorrower(borrower *models.Borrower) error {
	return s.borrowerRepo.Create(borrower)
}
