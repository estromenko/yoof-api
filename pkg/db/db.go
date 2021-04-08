package db

import (
	"io/ioutil"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Config struct {
	DSN             string        `json:"dsn" yaml:"dsn" mapstructure:"dsn"`
	MaxOpenConns    int           `json:"max_open_conns" yaml:"max_open_conns" mapstructure:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns" yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
	MigrationsPath  string        `json:"migrations_path" yaml:"migrations_path" mapstructure:"migrations_path"`
}

type Database struct {
	logger *zerolog.Logger
	config *Config
	db     *sqlx.DB
}

func New(logger *zerolog.Logger, config *Config) *Database {
	return &Database{
		config: config,
		logger: logger,
	}
}

func (d *Database) Open() error {
	conn, err := sqlx.Connect("pgx", d.config.DSN)
	if err != nil {
		return err
	}

	conn.SetMaxOpenConns(d.config.MaxOpenConns)
	conn.SetMaxIdleConns(d.config.MaxIdleConns)
	conn.SetConnMaxLifetime(d.config.ConnMaxLifetime)

	d.db = conn
	d.logger.Debug().Msg("Database connection opened successfully.")
	return nil
}

func (d *Database) Close() error {
	d.logger.Debug().Msg("Closing database connection.")
	if d.db == nil {
		return nil
	}
	return d.db.Close()
}

func (d *Database) Migrate() error {
	d.logger.Debug().Msg("Running migrations ...")
	files, err := ioutil.ReadDir(d.config.MigrationsPath)
	if err != nil {
		d.logger.Error().Msg(err.Error())
		return err
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(d.config.MigrationsPath + "/" + file.Name())

		if err != nil {
			d.logger.Error().Msg(err.Error())
			return err
		}

		if _, err := d.db.Exec(string(data)); err != nil {
			d.logger.Error().Msg(err.Error())
			return err
		}

		d.logger.Debug().Msg("- " + file.Name() + ": done.")
	}
	d.logger.Debug().Msg("Migrated successfully.")
	return nil
}

func (d *Database) DB() *sqlx.DB {
	return d.db
}
