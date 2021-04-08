package config_test

import (
	"testing"

	"github.com/estromenko/yoof-api/pkg/config"
	"github.com/stretchr/testify/assert"
)

type Config struct {
	Server string `mapstructure:"server"`
	DB     string `mapstructure:"db"`
	Logger string `mapstructure:"logger"`
}

func TestConfig(t *testing.T) {
	var conf Config
	assert.NoError(t, config.Load("../../configs/test.json", &conf))
	assert.Equal(t, "asd", conf.Server)
	assert.Equal(t, "asd", conf.DB)
	assert.Equal(t, "dsa", conf.Logger)
}
