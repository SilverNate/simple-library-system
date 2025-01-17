package services

import (
	"library-system/helper/times"
	"library-system/models"
	"library-system/repositories"
	"log"
	"time"
)

type historyService struct {
	historyRepo repositories.HistoryRepository
}

func NewHistoryService(repo repositories.HistoryRepository) HistoryService {
	return &historyService{historyRepo: repo}
}

func (s *historyService) AddHistoryRecord(request models.RequestHistoryBorrower) error {
	borrowerDate, err := time.Parse(times.TimeFormat, request.BorrowDate)
	if err != nil {
		log.Printf("Error in parsing: %v", err)
		return err
	}

	dueDate, err := time.Parse(times.TimeFormat, request.DueDate)
	if err != nil {
		log.Printf("Error in parsing: %v", err)
		return err
	}

	history := models.HistoryBorrower{
		BookID:     request.BookID,
		BorrowerID: request.BorrowerID,
		BorrowDate: borrowerDate.Local(),
		DueDate:    dueDate.Local(),
	}
	return s.historyRepo.Create(&history)
}

func (s *historyService) GetHistory(offset, pageSize int) ([]models.HistoryBorrower, error) {
	histories, err := s.historyRepo.GetHistory(offset, pageSize)
	if err != nil {
		return nil, err
	}

	for i, record := range histories {
		if record.ReturnDate == nil {
			continue
		}

		returnDate, _ := time.Parse(times.TimeFormat, record.ReturnDate.String())
		dueDate, _ := time.Parse(times.TimeFormat, record.DueDate.String())

		if returnDate.After(dueDate) {
			histories[i].Status = DefaultLate
		} else {
			histories[i].Status = DefaultOnTime
		}
		s.historyRepo.Update(&histories[i])
	}
	return histories, nil
}

func (s *historyService) UpdateReturnDate(id uint, returnDate string) error {
	history, err := s.historyRepo.GetHistoryById(id)
	if err != nil {
		log.Printf("Error get history by id: %v", err)
		return err
	}

	parsingReturnDate, err := time.Parse(times.TimeFormat, returnDate)
	if err != nil {
		log.Printf("Error in parsing return date: %v", err)
		return err
	}
	history.ReturnDate = &parsingReturnDate

	nextReturnDateParsed, err := time.Parse(times.TimeFormat, history.DueDate.Format(times.TimeFormat))
	if err != nil {
		log.Printf("Error in parsing next return date: %v", err)
		return err
	}

	if parsingReturnDate.After(nextReturnDateParsed) {
		history.Status = DefaultLate
	} else {
		history.Status = DefaultOnTime
	}

	return s.historyRepo.Update(history)
}

func (s *historyService) GetOverdueBooks(offset, pageSize int) ([]models.HistoryBorrower, error) {
	return s.historyRepo.FindOverdueBooks(offset, pageSize)
}

func (s *historyService) FindMostBorrowedBooks(offset, pageSize int) ([]models.Book, error) {
	return s.historyRepo.FindMostBorrowedBooks(offset, pageSize)
}
