package models

import (
	"time"
)

type HistoryBorrower struct {
	ID         uint `gorm:"primaryKey"`
	BookID     uint
	BorrowerID uint
	BorrowDate time.Time  `gorm:"not null;type:timestamp"`
	ReturnDate *time.Time `gorm:"default:NULL;type:timestamp"`
	DueDate    time.Time  `gorm:"not null;type:timestamp"`
	Status     string
}

type RequestHistoryBorrower struct {
	BookID     uint
	BorrowerID uint
	BorrowDate string `json:"borrowDate"`
	ReturnDate string `json:"returnDate"`
	DueDate    string `json:"nextReturnDate"`
	Status     string
}
