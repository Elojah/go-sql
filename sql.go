package sql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Service wraps a SQL client.
type Service struct {
	*sql.DB
}

// Dial connects SQL server.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	var err error

	s.DB, err = sql.Open("postgres", info)
	if err != nil {
		return err
	}

	return s.DB.Ping()
}

func (s *Service) Close(ctx context.Context) error {
	if s.DB != nil {
		return s.DB.Close()
	}

	return nil
}
