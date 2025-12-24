package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableUrls, downCreateTableUrls)
}

func upCreateTableUrls(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE url (
						id SERIAL PRIMARY KEY,
						short_code VARCHAR(10) UNIQUE NOT NULL,
						original_url TEXT NOT NULL,
						created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
						expires_at TIMESTAMP WITH TIME ZONE,
						user_id INTEGER, -- для будущей аутентификации
						visits INTEGER DEFAULT 0,
						last_visited_at TIMESTAMP WITH TIME ZONE
					);`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`CREATE INDEX idx_urls_short_code ON urls(short_code);`)

	if err != nil {
		return err
	}

	_, err = tx.Exec(`CREATE INDEX idx_urls_user_id ON urls(user_id);`)

	if err != nil {
		return err
	}
	return nil
}

func downCreateTableUrls(ctx context.Context, tx *sql.Tx) error {

	return nil
}
