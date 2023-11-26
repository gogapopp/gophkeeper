package repository

import (
	"errors"

	"github.com/jackc/pgconn"
)

var (
	PgErr                *pgconn.PgError
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotExists     = errors.New("user not exists")
)
