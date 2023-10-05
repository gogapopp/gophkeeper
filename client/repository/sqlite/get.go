package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/gogapopp/gophkeeper/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TODO:
const layout = "2006-01-02 15:04:05.999999-07:00"

func (r *Repository) GetUniqueKeys(ctx context.Context, userID int) (map[string][]string, error) {
	const op = "sqlite.get.GetUniqueKeys"
	tables := []string{"textdata", "binarydata", "carddata"}
	keys := make(map[string][]string)
	for _, table := range tables {
		rows, err := r.db.QueryContext(ctx, fmt.Sprintf("SELECT unique_key FROM %s WHERE user_id=?1", table), userID)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
		defer rows.Close()
		var uniqueKeys []string
		for rows.Next() {
			var uniqueKey string
			if err := rows.Scan(&uniqueKey); err != nil {
				return nil, fmt.Errorf("%s: %s", op, err)
			}
			uniqueKeys = append(uniqueKeys, uniqueKey)
		}
		keys[table] = uniqueKeys
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
	}
	return keys, nil
}

func (r *Repository) GetTextData(ctx context.Context, uniqueKey int) (models.TextData, error) {
	const op = "sqlite.get.GetTextData"
	const query = "SELECT unique_key, text_data, uploaded_at, metainfo FROM textdata WHERE unique_key=?1"
	var textdata models.TextData
	var uploadedAtStr string
	rows, err := r.db.QueryContext(ctx, query, uniqueKey)
	if err != nil {
		return models.TextData{}, fmt.Errorf("%s: %s", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&textdata.UniqueKey, &textdata.TextData, &uploadedAtStr, &textdata.Metainfo)
		if err != nil {
			return models.TextData{}, fmt.Errorf("%s: %s", op, err)
		}
		uploadedAt, err := time.Parse(layout, uploadedAtStr)
		if err != nil {
			return models.TextData{}, fmt.Errorf("%s: %s", op, err)
		}
		textdata.UploadedAt = timestamppb.New(uploadedAt)
	}
	return textdata, nil
}

func (r *Repository) GetBinaryData(ctx context.Context, uniqueKey int) (models.BinaryData, error) {
	const op = "sqlite.get.GetBinaryData"
	const query = "SELECT binary_data, metainfo, uploaded_at FROM binarydata WHERE unique_key=?1"
	var binarydata models.BinaryData
	var uploadedAtStr string
	rows, err := r.db.QueryContext(ctx, query, uniqueKey)
	if err != nil {
		return models.BinaryData{}, fmt.Errorf("%s: %s", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&binarydata.BinaryData, &binarydata.Metainfo, &uploadedAtStr)
		if err != nil {
			return models.BinaryData{}, fmt.Errorf("%s: %s", op, err)
		}
		uploadedAt, err := time.Parse(layout, uploadedAtStr)
		if err != nil {
			return models.BinaryData{}, fmt.Errorf("%s: %s", op, err)
		}
		binarydata.UploadedAt = timestamppb.New(uploadedAt)
	}
	return binarydata, nil
}

func (r *Repository) GetCardData(ctx context.Context, uniqueKey int) (models.CardData, error) {
	const op = "sqlite.get.GetCardData"
	const query = "SELECT unique_key, card_number, card_name, card_date, cvv, uploaded_at, metainfo FROM carddata WHERE unique_key=?1"
	var carddata models.CardData
	var uploadedAtStr string
	rows, err := r.db.QueryContext(ctx, query, uniqueKey)
	if err != nil {
		return models.CardData{}, fmt.Errorf("%s: %s", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&carddata.UniqueKey, &carddata.CardNumberData, &carddata.CardNameData, &carddata.CardDateData, &carddata.CvvData, &uploadedAtStr, &carddata.Metainfo)
		if err != nil {
			return models.CardData{}, fmt.Errorf("%s: %s", op, err)
		}
		uploadedAt, err := time.Parse(layout, uploadedAtStr)
		if err != nil {
			return models.CardData{}, fmt.Errorf("%s: %s", op, err)
		}
		carddata.UploadedAt = timestamppb.New(uploadedAt)
	}
	return carddata, nil
}

func (r *Repository) GetDatas(ctx context.Context, table string) (map[int]string, error) {
	const op = "sqlite.get.GetDatas"
	query := fmt.Sprintf("SELECT unique_key, uploaded_at FROM %s", table)
	datas := make(map[int]string)
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			uniqueKey  int
			uploadedAt string
		)
		if err := rows.Scan(&uniqueKey, &uploadedAt); err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
		datas[uniqueKey] = uploadedAt
	}
	return datas, nil
}
