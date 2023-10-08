package service

import "github.com/gogapopp/gophkeeper/client/repository/sqlite"

type SaveService struct {
	store Storager
}

type HashService struct {
}

type GetService struct {
	get Getter
}

// NewSaveService возвращает экземпляр структуры сервиса для сохранения данных
func NewSaveService(repository *sqlite.Repository) *SaveService {
	return &SaveService{store: repository}
}

// NewHashService возвращает экземпляр структуры сервиса для хеширования данных
func NewHashService() *HashService {
	return &HashService{}
}

// NewGetService возвращает эксземляр структуры сервиса для получения данных
func NewGetService(repository *sqlite.Repository) *GetService {
	return &GetService{get: repository}
}
