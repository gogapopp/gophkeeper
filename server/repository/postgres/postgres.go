package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Repository struct {
	db *sql.DB
}

// NewRepo подключается к БД и создаёт sql таблицы
func NewRepo(serverDBdsn string) (*Repository, *sql.DB, error) {
	const op = "postgres.postgresql.NewRepo"
	db, err := sql.Open("pgx", serverDBdsn)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %s", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %s", op, err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		login VARCHAR(256) NOT NULL UNIQUE,
		password VARCHAR(256) NOT NULL,
		user_phrase VARCHAR(256) NOT NULL,
		last_update_at TIMESTAMPTZ
	);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_login ON users(login);

	CREATE TABLE IF NOT EXISTS textdata (
		text_data_id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		unique_key VARCHAR(128),
		text_data BYTEA NOT NULL,
		uploaded_at TIMESTAMPTZ NOT NULL,
		metainfo BYTEA,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS binarydata (
		binary_data_id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		unique_key VARCHAR(128),
		binary_data BYTEA NOT NULL,
		uploaded_at TIMESTAMPTZ NOT NULL,
		metainfo BYTEA,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS carddata (
		card_data_id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		unique_key VARCHAR(128),
		card_number BYTEA NOT NULL,
		card_name BYTEA NOT NULL,
		card_date BYTEA NOT NULL,
		cvv BYTEA NOT NULL,
		uploaded_at TIMESTAMPTZ NOT NULL,
		metainfo BYTEA,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %s", op, err)
	}

	return &Repository{db: db}, db, nil
}
