package service

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

func (s *SaveService) AddTextData(ctx context.Context, textdata models.TextData) error {
	const op = "service.saver.AddTextData"
	err := s.store.AddTextData(ctx, textdata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (s *SaveService) AddBinaryData(ctx context.Context, binarydata models.BinaryData) error {
	const op = "service.saver.AddBinaryData"
	err := s.store.AddBinaryData(ctx, binarydata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (s *SaveService) AddCardData(ctx context.Context, carddata models.CardData) error {
	const op = "service.saver.AddCardData"
	err := s.store.AddCardData(ctx, carddata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}
