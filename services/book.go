package services

import (
	"library-system/models"
	"library-system/repositories"
)

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{bookRepo: repo}
}

func (s *bookService) CreateBook(book *models.Book) error {
	return s.bookRepo.Create(book)
}
