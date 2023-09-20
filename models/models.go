package models

import "time"

type (
	// User
	User struct {
		Login      string
		Password   string
		UploadedAt time.Time
	}

	BinaryData struct {
		UserID     int64
		UniqueKey  string
		BinaryData []byte
		UploadedAt string
		Metainfo   []byte
	}

	TextData struct {
		UserID     int64
		UniqueKey  string
		TextData   []byte
		UploadedAt string
		Metainfo   []byte
	}

	CardData struct {
		UserID         int64
		UniqueKey      string
		CardNumberData []byte
		CardNameData   []byte
		CardDateData   []byte
		CvvData        []byte
		UploadedAt     string
		Metainfo       []byte
	}
)
