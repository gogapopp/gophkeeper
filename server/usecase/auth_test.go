package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gogapopp/gophkeeper/internal/hasher"
	"github.com/gogapopp/gophkeeper/models"
	"github.com/gogapopp/gophkeeper/server/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	testcases := []struct {
		name    string
		user    models.User
		mockErr error
		wantErr bool
	}{
		{
			name: "Success",
			user: models.User{
				Login:      "login",
				Password:   "password",
				UserPhrase: "phrase",
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Error",
			user: models.User{
				Login:      "login",
				Password:   "",
				UserPhrase: "phrase",
			},
			mockErr: errors.New("error"),
			wantErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockAuth := new(mocks.MockRepo)
			ctx := context.TODO()
			hashedPassword := hasher.GenerateHash(tc.user.Password)
			hashedPhrase := hasher.GenerateHash(tc.user.UserPhrase)
			mockAuth.On("Register", ctx, models.User{
				Login:      tc.user.Login,
				Password:   hashedPassword,
				UserPhrase: hashedPhrase,
			}).Return(tc.mockErr)
			authUsecase := AuthUsecase{auth: mockAuth}
			err := authUsecase.Register(ctx, tc.user)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockAuth.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {
	testcases := []struct {
		name    string
		user    models.User
		mockErr error
		wantErr bool
	}{
		{
			name: "Success",
			user: models.User{
				Login:      "login",
				Password:   "password",
				UserPhrase: "phrase",
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Error",
			user: models.User{
				Login:      "login",
				Password:   "",
				UserPhrase: "phrase",
			},
			mockErr: errors.New("error"),
			wantErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockAuth := new(mocks.MockRepo)
			ctx := context.TODO()
			hashedPassword := hasher.GenerateHash(tc.user.Password)
			hashedPhrase := hasher.GenerateHash(tc.user.UserPhrase)
			if tc.wantErr {
				mockAuth.On("Login", ctx, models.User{
					Login:      tc.user.Login,
					Password:   hashedPassword,
					UserPhrase: hashedPhrase,
					UploadedAt: time.Now(),
				}).Return("1", tc.mockErr)
			} else {
				mockAuth.On("Login", ctx, models.User{
					Login:      tc.user.Login,
					Password:   hashedPassword,
					UserPhrase: hashedPhrase,
					UploadedAt: time.Now(),
				}).Return("1", nil)
			}
			authUsecase := AuthUsecase{auth: mockAuth}
			token, err := authUsecase.Login(ctx, tc.user)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, "1")
			}

			mockAuth.AssertExpectations(t)
		})
	}
}
