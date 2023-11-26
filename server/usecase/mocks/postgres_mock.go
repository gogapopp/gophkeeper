package mocks

import (
	"context"

	"github.com/gogapopp/gophkeeper/models"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Register(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepo) Login(ctx context.Context, user models.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *MockRepo) GetDatas(uniqueKeys map[string][]string) (models.SyncData, error) {
	args := m.Called(uniqueKeys)
	return args.Get(0).(models.SyncData), args.Error(1)
}

func (m *MockRepo) AddTextData(ctx context.Context, textdata models.TextData) error {
	args := m.Called(ctx, textdata)
	return args.Error(0)
}

func (m *MockRepo) AddBinaryData(ctx context.Context, binarydata models.BinaryData) error {
	args := m.Called(ctx, binarydata)
	return args.Error(0)
}

func (m *MockRepo) AddCardData(ctx context.Context, carddata models.CardData) error {
	args := m.Called(ctx, carddata)
	return args.Error(0)
}
