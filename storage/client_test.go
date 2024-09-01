package storage

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func NewDBClientWithMockPetsRepo(t *testing.T) (*DBClient, *MockPetsRepository) {
	ctrl := gomock.NewController(t)
	petsStore := NewMockPetsRepository(ctrl)
	return &DBClient{dbSession: nil, petsStore: petsStore}, petsStore
}
