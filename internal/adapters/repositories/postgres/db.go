package postgres

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alandavd/empomio/internal/config"
)

// NewPostgresDB creates a new PostgreSQL database instance.
func NewPostgresDB(ctx context.Context, config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		config.PostgreSQLUser(),
		config.PostgreSQLPassword(),
		config.PostgreSQLName(),
		config.PostgreSQLHost(),
		config.PostgreSQLPort(),
		config.PostgreSQLSSLMode())
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
		return db,nil
}
