package usecase

import "github.com/gogapopp/gophkeeper/server/repository/postgres"

type (
	AuthUsecase struct {
		auth Auth
	}

	StorageUsecase struct {
		store Storager
	}

	GetUsecase struct {
		get Getter
	}
)

func NewAuthUsecase(repository *postgres.Repository) *AuthUsecase {
	return &AuthUsecase{auth: repository}
}

func NewStorageUsecase(repository *postgres.Repository) *StorageUsecase {
	return &StorageUsecase{store: repository}
}

func NewGetUsecase(repository *postgres.Repository) *GetUsecase {
	return &GetUsecase{get: repository}
}
