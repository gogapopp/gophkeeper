package service

import (
	"context"
	"fmt"

	pb "github.com/gogapopp/gophkeeper/proto"

	"github.com/gogapopp/gophkeeper/models"
)

// интерфейс взаимодействия с БД для сохранения файлов
type Storager interface {
	AddTextData(ctx context.Context, textdata models.TextData) error
	AddBinaryData(ctx context.Context, binarydata models.BinaryData) error
	AddCardData(ctx context.Context, carddata models.CardData) error
	SaveDatas(ctx context.Context, syncdata models.SyncData) error
}

// AddTextData возвращает текстовые данные
func (s *SaveService) AddTextData(ctx context.Context, textdata models.TextData) error {
	const op = "service.saver.AddTextData"
	err := s.store.AddTextData(ctx, textdata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

// AddBinaryData возвращает бинарные данные
func (s *SaveService) AddBinaryData(ctx context.Context, binarydata models.BinaryData) error {
	const op = "service.saver.AddBinaryData"
	err := s.store.AddBinaryData(ctx, binarydata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

// AddCardData возвращает данные карты
func (s *SaveService) AddCardData(ctx context.Context, carddata models.CardData) error {
	const op = "service.saver.AddCardData"
	err := s.store.AddCardData(ctx, carddata)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

// SaveDatas получаем данные из запроса в модели данных для сохранения в БД
func (s *SaveService) SaveDatas(ctx context.Context, syncdata *pb.SyncResponse) error {
	const op = "service.saver.SyncData"
	if syncdata == nil {
		return nil
	}
	var textData []models.TextData
	for _, td := range syncdata.TextData {
		textData = append(textData, models.TextData{
			TextData:   td.TextData,
			Metainfo:   td.Metainfo,
			UserID:     td.UserID,
			UniqueKey:  td.UniqueKey,
			UploadedAt: td.UploadedAt,
		})
	}
	var binaryData []models.BinaryData
	for _, bd := range syncdata.BinaryData {
		binaryData = append(binaryData, models.BinaryData{
			BinaryData: bd.BinaryData,
			Metainfo:   bd.Metainfo,
			UserID:     bd.UserID,
			UniqueKey:  bd.UniqueKey,
			UploadedAt: bd.UploadedAt,
		})
	}
	var cardData []models.CardData
	for _, cd := range syncdata.CardData {
		cardData = append(cardData, models.CardData{
			CardNumberData: cd.CardNumberData,
			CardNameData:   cd.CardNameData,
			CardDateData:   cd.CardDateData,
			CvvData:        cd.CvvData,
			Metainfo:       cd.Metainfo,
			UserID:         cd.UserID,
			UniqueKey:      cd.UniqueKey,
			UploadedAt:     cd.UploadedAt,
		})
	}
	preparedSyncdata := models.SyncData{
		TextData:   textData,
		BinaryData: binaryData,
		CardData:   cardData,
	}
	err := s.store.SaveDatas(ctx, preparedSyncdata)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
