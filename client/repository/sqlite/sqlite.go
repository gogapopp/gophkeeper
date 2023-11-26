package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

// NewRepo подключается к БД и создаёт sql таблицы
func NewRepo(clientDBdsn string) (*Repository, *sql.DB, error) {
	const op = "sqlite.sqlite.NewRepo"
	db, err := sql.Open("sqlite3", clientDBdsn)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = db.Exec(`
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
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Repository{db: db}, db, nil
}
