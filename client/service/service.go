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

func NewSaveService(repository *sqlite.Repository) *SaveService {
	return &SaveService{store: repository}
}

func NewHashService() *HashService {
	return &HashService{}
}

func NewGetService(repository *sqlite.Repository) *GetService {
	return &GetService{get: repository}
}
