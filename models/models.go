package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	// модель пользователя
	User struct {
		Login      string
		Password   string
		UserPhrase string
		UploadedAt time.Time
	}
	// модель бинарный данных
	BinaryData struct {
		UserID     int64
		UniqueKey  string
		BinaryData []byte
		UploadedAt *timestamppb.Timestamp
		Metainfo   []byte
	}
	// модель текстовых данных
	TextData struct {
		UserID     int64
		UniqueKey  string
		TextData   []byte
		UploadedAt *timestamppb.Timestamp
		Metainfo   []byte
	}
	// модель данных карты
	CardData struct {
		UserID         int64
		UniqueKey      string
		CardNumberData []byte
		CardNameData   []byte
		CardDateData   []byte
		CvvData        []byte
		UploadedAt     *timestamppb.Timestamp
		Metainfo       []byte
	}
	// модель содержащая все типы данных
	SyncData struct {
		TextData   []TextData
		BinaryData []BinaryData
		CardData   []CardData
	}
)
