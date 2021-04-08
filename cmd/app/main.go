package main

import (
	"flag"

	"github.com/estromenko/yoof-api/internal/app"
	_ "github.com/estromenko/yoof-api/internal/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "", "config path")
}

// @title YooF API Documentation
// @version 1.0
// @description API documentation

func main() {
	flag.Parse()
	app.Run(configPath)
}
