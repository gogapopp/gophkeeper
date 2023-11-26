package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogapopp/gophkeeper/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetDatas получает все данные пользователя по уникальным ключам
func (r *Repository) GetDatas(uniqueKeys map[string][]string) (models.SyncData, error) {
	const op = "postgres.get.GetDatas"
	var syncData models.SyncData
	for tableName, keys := range uniqueKeys {
		// если у пользователя на клиенте нет было сохранёных данных
		if len(keys) == 0 {
			keys = append(keys, "00000000")
		}
		// плейсхолдер для сохранения сразу отправления большого кол-ва ключей в запрос
		placeholders := make([]string, len(keys))
		for i := range keys {
			placeholders[i] = "$" + strconv.Itoa(i+1)
		}
		// получаем нужный нам запрос
		var query string
		switch tableName {
		case "textdata":
			query = fmt.Sprintf("SELECT user_id, unique_key, text_data, uploaded_at, metainfo FROM %s WHERE unique_key NOT IN (%s)", tableName, strings.Join(placeholders, ","))
		case "binarydata":
			query = fmt.Sprintf("SELECT user_id, unique_key, binary_data, uploaded_at, metainfo FROM %s WHERE unique_key NOT IN (%s)", tableName, strings.Join(placeholders, ","))
		case "carddata":
			query = fmt.Sprintf("SELECT user_id, unique_key, card_number, card_name, card_date, cvv, uploaded_at, metainfo FROM %s WHERE unique_key NOT IN (%s)", tableName, strings.Join(placeholders, ","))
		default:
			return models.SyncData{}, fmt.Errorf("%s: table is not exists", op)
		}
		// уникальные ключи для типа данных
		args := make([]interface{}, len(keys))
		for i, key := range keys {
			args[i] = key
		}
		// отправляем запрос
		rows, err := r.db.Query(query, args...)
		if err != nil {
			return models.SyncData{}, fmt.Errorf("%s: %w", op, err)
		}
		defer rows.Close()
		// получаем данные из БД для каждого типа данных
		switch tableName {
		case "textdata":
			for rows.Next() {
				var td models.TextData
				var uploadedAt time.Time
				err := rows.Scan(&td.UserID,
					&td.UniqueKey,
					&td.TextData,
					&uploadedAt,
					&td.Metainfo)
				if err != nil {
					if err == sql.ErrNoRows {
						return models.SyncData{}, err
					}
					return models.SyncData{}, fmt.Errorf("%s: %w", op, err)
				}
				td.UploadedAt = timestamppb.New(uploadedAt)
				syncData.TextData = append(syncData.TextData, td)
			}
		case "binarydata":
			for rows.Next() {
				var bd models.BinaryData
				var uploadedAt time.Time
				err := rows.Scan(&bd.UserID,
					&bd.UniqueKey,
					&bd.BinaryData,
					&uploadedAt,
					&bd.Metainfo)
				if err != nil {
					if err == sql.ErrNoRows {
						return models.SyncData{}, err
					}
					return models.SyncData{}, fmt.Errorf("%s: %w", op, err)
				}
				bd.UploadedAt = timestamppb.New(uploadedAt)
				syncData.BinaryData = append(syncData.BinaryData, bd)
			}
		case "carddata":
			for rows.Next() {
				var cd models.CardData
				var uploadedAt time.Time
				err := rows.Scan(&cd.UserID,
					&cd.UniqueKey,
					&cd.CardNumberData,
					&cd.CardNameData,
					&cd.CardDateData,
					&cd.CvvData,
					&uploadedAt,
					&cd.Metainfo)
				if err != nil {
					if err == sql.ErrNoRows {
						return models.SyncData{}, err
					}
					return models.SyncData{}, fmt.Errorf("%s: %w", op, err)
				}
				cd.UploadedAt = timestamppb.New(uploadedAt)
				syncData.CardData = append(syncData.CardData, cd)
			}
		default:
			return models.SyncData{}, fmt.Errorf("%s: table is not exists", op)
		}
		if err := rows.Err(); err != nil {
			if err == sql.ErrNoRows {
				return models.SyncData{}, err
			}
			return models.SyncData{}, fmt.Errorf("%s: %w", op, err)
		}
	}
	return syncData, nil
}
