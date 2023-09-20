package usecase

import (
	"context"
	"fmt"

	"github.com/gogapopp/gophkeeper/models"
)

type Storager interface {
	AddTextData(ctx context.Context, textdata models.TextData) error
	AddBinaryData(ctx context.Context, binarydata models.BinaryData) error
	AddCardData(ctx context.Context, carddata models.CardData) error
}

func (su *StorageUsecase) AddTextData(ctx context.Context, textdata models.TextData) error {
	const op = "usecase.store.AddTextData"
	err := su.store.AddTextData(ctx, textdata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (su *StorageUsecase) AddBinaryData(ctx context.Context, binarydata models.BinaryData) error {
	const op = "usecase.store.AddBinaryData"
	err := su.store.AddBinaryData(ctx, binarydata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (su *StorageUsecase) AddCardData(ctx context.Context, carddata models.CardData) error {
	const op = "usecase.store.AddCardData"
	err := su.store.AddCardData(ctx, carddata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}
