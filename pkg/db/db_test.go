package db_test

import (
	"testing"
	"time"

	"github.com/estromenko/yoof-api/pkg/db"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	tests := []struct {
		name                  string
		config                *db.Config
		isConfigPathError     bool
		isMigrationsPathError bool
	}{
		{
			name: "Right config",
			config: &db.Config{
				DSN:             "user=postgres password=1234 host=localhost port=5432 sslmode=disable",
				MaxOpenConns:    25,
				MaxIdleConns:    25,
				ConnMaxLifetime: time.Minute * 5,
				MigrationsPath:  "../../migrations",
			},
			isConfigPathError:     false,
			isMigrationsPathError: false,
		},
		{
			name: "Bad config",
			config: &db.Config{
				DSN:             "asdasdasd asdasdasd asdasd",
				MaxOpenConns:    25,
				MaxIdleConns:    25,
				ConnMaxLifetime: time.Minute * 5,
				MigrationsPath:  "migrations",
			},
			isConfigPathError:     true,
			isMigrationsPathError: false,
		},
		{
			name: "Bad migrations path",
			config: &db.Config{
				DSN:             "user=postgres password=1234 host=localhost port=5432 sslmode=disable",
				MaxOpenConns:    25,
				MaxIdleConns:    25,
				ConnMaxLifetime: time.Minute * 5,
				MigrationsPath:  "migration",
			},
			isConfigPathError:     false,
			isMigrationsPathError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := db.New(&zerolog.Logger{}, tt.config)
			assert.Equal(t, tt.isConfigPathError, database.Open() != nil, "Error opening database connection")
			if !tt.isConfigPathError {
				assert.Equal(t, tt.isMigrationsPathError, database.Migrate() != nil, "Error running migrations")
				assert.NoError(t, database.Close(), "Error closing database connection")
			}
		})
	}
}
