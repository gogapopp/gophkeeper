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

// NewAuthUsecase возвращает экземпляр структуры сервиса аутентификаци
func NewAuthUsecase(repository *postgres.Repository) *AuthUsecase {
	return &AuthUsecase{auth: repository}
}

// NewStorageUsecase возвращает экземпляр структуры сервиса сохранения
func NewStorageUsecase(repository *postgres.Repository) *StorageUsecase {
	return &StorageUsecase{store: repository}
}

// NewGetUsecase возвращает жкземпляр структуры для получения (синхронизации) данных
func NewGetUsecase(repository *postgres.Repository) *GetUsecase {
	return &GetUsecase{get: repository}
}
