package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/gogapopp/gophkeeper/models"
	"github.com/gogapopp/gophkeeper/server/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddTextData(t *testing.T) {
	testcases := []struct {
		name    string
		data    models.TextData
		mockErr error
		wantErr bool
	}{
		{
			name: "Success",
			data: models.TextData{
				TextData: []byte("111111111111111"),
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Error",
			data: models.TextData{
				TextData: []byte("111111111111111"),
			},
			mockErr: errors.New("error"),
			wantErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := new(mocks.MockRepo)
			ctx := context.TODO()
			mockStorage.On("AddTextData", ctx, tc.data).Return(tc.mockErr)
			storageUsecase := StorageUsecase{store: mockStorage}
			err := storageUsecase.AddTextData(ctx, tc.data)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockStorage.AssertExpectations(t)
		})
	}
}

func TestAddBinaryData(t *testing.T) {
	testcases := []struct {
		name    string
		data    models.BinaryData
		mockErr error
		wantErr bool
	}{
		{
			name: "Success",
			data: models.BinaryData{
				BinaryData: []byte("1111"),
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Error",
			data: models.BinaryData{
				BinaryData: []byte("1111"),
			},
			mockErr: errors.New("error"),
			wantErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := new(mocks.MockRepo)
			ctx := context.TODO()
			mockStorage.On("AddBinaryData", ctx, tc.data).Return(tc.mockErr)
			storageUsecase := StorageUsecase{store: mockStorage}
			err := storageUsecase.AddBinaryData(ctx, tc.data)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockStorage.AssertExpectations(t)
		})
	}
}

func TestAddCardData(t *testing.T) {
	testcases := []struct {
		name    string
		data    models.CardData
		mockErr error
		wantErr bool
	}{
		{
			name: "Success",
			data: models.CardData{
				CardNumberData: []byte("11111111111"),
				CardDateData:   []byte("11111111111"),
				CardNameData:   []byte("11111111111"),
				CvvData:        []byte("11111111111"),
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Error",
			data: models.CardData{
				CardNumberData: []byte("11111111111"),
				CardDateData:   []byte("11111111111"),
				CardNameData:   []byte("11111111111"),
				CvvData:        []byte("11111111111"),
			},
			mockErr: errors.New("error"),
			wantErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := new(mocks.MockRepo)
			ctx := context.TODO()
			mockStorage.On("AddCardData", ctx, tc.data).Return(tc.mockErr)
			storageUsecase := StorageUsecase{store: mockStorage}
			err := storageUsecase.AddCardData(ctx, tc.data)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockStorage.AssertExpectations(t)
		})
	}
}
