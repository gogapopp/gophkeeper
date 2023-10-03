package service

import "github.com/gogapopp/gophkeeper/client/repository/sqlite"

type SaveService struct {
	store Storager
}

type HashService struct {
}

func NewSaveService(repository *sqlite.Repository) *SaveService {
	return &SaveService{store: repository}
}

func NewHashService() *HashService {
	return &HashService{}
}
