package app

import (
	"log"

	"github.com/estromenko/yoof-api/internal/server"
	"github.com/estromenko/yoof-api/pkg/config"
	"github.com/estromenko/yoof-api/pkg/db"
	"github.com/estromenko/yoof-api/pkg/logger"
)

type Config struct {
	Logger   logger.Config `json:"logger" yaml:"logger"`
	Database db.Config     `json:"database" yaml:"database"`
	Server   server.Config `json:"server" yaml:"server"`
}

func Run(configPath string) {
	if configPath == "" {
		log.Fatal("config path is not setted.")
	}

	var conf Config
	err := config.Load(configPath, &conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := logger.New(&conf.Logger)
	database := db.New(logger, &conf.Database)
	server := server.New(database, logger, &conf.Server)

	if err := database.Open(); err != nil {
		logger.Fatal().Msg(err.Error())
	}

	if err := database.Migrate(); err != nil {
		return
	}

	logger.Fatal().Msg(server.Run().Error())
}
