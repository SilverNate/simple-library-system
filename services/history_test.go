package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"library-system/models"
	"testing"
)

// Mock Repository
type MockHistoryRepository struct {
	mock.Mock
}

func (m *MockHistoryRepository) Create(history *models.HistoryBorrower) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockHistoryRepository) Update(history *models.HistoryBorrower) error {
	args := m.Called(history)
	return args.Error(1)
}

func (m *MockHistoryRepository) GetHistory(offset, pageSize int) ([]models.HistoryBorrower, error) {
	args := m.Called(offset, pageSize)
	return args.Get(0).([]models.HistoryBorrower), args.Error(1)
}

func (m *MockHistoryRepository) GetHistoryById(id uint) (*models.HistoryBorrower, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.HistoryBorrower), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockHistoryRepository) FindOverdueBooks(offset, pageSize int) ([]models.HistoryBorrower, error) {
	args := m.Called(offset, pageSize)
	return args.Get(0).([]models.HistoryBorrower), args.Error(1)
}

func (m *MockHistoryRepository) FindMostBorrowedBooks(offset, pageSize int) ([]models.Book, error) {
	args := m.Called(offset, pageSize)
	return args.Get(0).([]models.Book), args.Error(1)
}

func TestGetOverdueBooks(t *testing.T) {
	mockRepo := new(MockHistoryRepository)
	historyService := NewHistoryService(mockRepo)

	mockBooks := []models.HistoryBorrower{
		{ID: 1, BookID: 1, BorrowerID: 2, Status: "late"},
	}
	mockRepo.On("FindOverdueBooksPaginated", 1, 10).Return(mockBooks, nil)

	result, err := historyService.GetOverdueBooks(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].Status, "late")
	mockRepo.AssertExpectations(t)
}
