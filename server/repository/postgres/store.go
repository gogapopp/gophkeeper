package postgres

import (
	"context"
	"fmt"

	"github.com/gogapopp/gophkeeper/models"
)

func (r *Repository) AddTextData(ctx context.Context, textdata models.TextData) error {
	const op = "postgres.store.AddTextData"
	const query = "INSERT INTO textdata (user_id, unique_key, text_data, uploaded_at, metainfo) values ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, textdata.UserID, textdata.UniqueKey, textdata.TextData, textdata.UploadedAt, textdata.Metainfo)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (r *Repository) AddBinaryData(ctx context.Context, binarydata models.BinaryData) error {
	const op = "postgres.store.AddBinaryData"
	const query = "INSERT INTO binarydata (user_id, unique_key, binary_data, uploaded_at, metainfo) values ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, binarydata.UserID, binarydata.UniqueKey, binarydata.BinaryData, binarydata.UploadedAt, binarydata.Metainfo)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (r *Repository) AddCardData(ctx context.Context, carddata models.CardData) error {
	const op = "postgres.store.AddCardData"
	const query = "INSERT INTO carddata (user_id, unique_key, card_number, card_name, card_date, cvv, uploaded_at, metainfo) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := r.db.ExecContext(ctx, query, carddata.UserID, carddata.UniqueKey, carddata.CardNumberData, carddata.CardNameData, carddata.CardDateData, carddata.CvvData, carddata.UploadedAt, carddata.Metainfo)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}
