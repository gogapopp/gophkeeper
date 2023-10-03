package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

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
		UploadedAt *timestamppb.Timestamp
		Metainfo   []byte
	}

	TextData struct {
		UserID     int64
		UniqueKey  string
		TextData   []byte
		UploadedAt *timestamppb.Timestamp
		Metainfo   []byte
	}

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
)
