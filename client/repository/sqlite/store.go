package sqlite

import (
	"context"
	"fmt"

	"github.com/gogapopp/gophkeeper/models"
)

const (
	textDataQuery   = "INSERT INTO textdata (user_id, unique_key, text_data, uploaded_at, metainfo) values (?1, ?2, ?3, ?4, ?5)"
	binaryDataQuery = "INSERT INTO binarydata (user_id, unique_key, binary_data, uploaded_at, metainfo) values (?1, ?2, ?3, ?4, ?5)"
	cardDataQuery   = "INSERT INTO carddata (user_id, unique_key, card_number, card_name, card_date, cvv, uploaded_at, metainfo) values (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8)"
)

func (r *Repository) AddTextData(ctx context.Context, textdata models.TextData) error {
	const op = "sqlite.store.AddTextData"
	_, err := r.db.ExecContext(ctx, textDataQuery, textdata.UserID, textdata.UniqueKey, textdata.TextData, textdata.UploadedAt.AsTime(), textdata.Metainfo)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (r *Repository) AddBinaryData(ctx context.Context, binarydata models.BinaryData) error {
	const op = "sqlite.store.AddBinaryData"
	_, err := r.db.ExecContext(ctx, binaryDataQuery, binarydata.UserID, binarydata.UniqueKey, binarydata.BinaryData, binarydata.UploadedAt.AsTime(), binarydata.Metainfo)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (r *Repository) AddCardData(ctx context.Context, carddata models.CardData) error {
	const op = "sqlite.store.AddCardData"
	_, err := r.db.ExecContext(ctx, binaryDataQuery,
		carddata.UserID, carddata.UniqueKey, carddata.CardNumberData, carddata.CardNameData, carddata.CardDateData, carddata.CvvData, carddata.UploadedAt.AsTime(), carddata.Metainfo)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (r *Repository) SaveDatas(ctx context.Context, syncdata models.SyncData) error {
	const op = "sqlite.store.SyncData"
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	for _, textdata := range syncdata.TextData {
		_, err := tx.ExecContext(ctx, textDataQuery, textdata.UserID, textdata.UniqueKey, textdata.TextData, textdata.UploadedAt.AsTime(), textdata.Metainfo)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("%s: %s", op, err)
		}
	}
	for _, binarydata := range syncdata.BinaryData {
		_, err := tx.ExecContext(ctx, binaryDataQuery, binarydata.UserID, binarydata.UniqueKey, binarydata.BinaryData, binarydata.UploadedAt.AsTime(), binarydata.Metainfo)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("%s: %s", op, err)
		}
	}
	for _, carddata := range syncdata.CardData {
		_, err := tx.ExecContext(ctx, binaryDataQuery,
			carddata.UserID, carddata.UniqueKey, carddata.CardNumberData, carddata.CardNameData, carddata.CardDateData, carddata.CvvData, carddata.UploadedAt.AsTime(), carddata.Metainfo)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("%s: %s", op, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}
