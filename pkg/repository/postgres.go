package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Hymiside/it-market-api/pkg/models"
)

func NewPostgresDB(ctx context.Context, c models.ConfigRepository) (*sql.DB, error) {
	connRequest := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
	db, err := sql.Open("postgres", connRequest)
	if err != nil {
		return nil, fmt.Errorf("error to connection db: %v", err)
	}
	go func(ctx context.Context) {
		<-ctx.Done()
		db.Close()
	}(ctx)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("connection test error: %w", err)
	}
	return db, nil
}
