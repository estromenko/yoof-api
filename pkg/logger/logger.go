package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type LoggerWriter struct {
	w     zerolog.LevelWriter
	level zerolog.Level
}

func (w *LoggerWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w *LoggerWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level == w.level {
		return w.w.WriteLevel(level, p)
	}
	return len(p), nil
}

type Config struct {
	LogsPath string `json:"logs_path" yaml:"logs_path" mapstructure:"logs_path"`
}

func New(config *Config) *zerolog.Logger {
	fInfo, _ := os.OpenFile(config.LogsPath+"/info.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	fError, _ := os.OpenFile(config.LogsPath+"/error.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	infoWriter := zerolog.MultiLevelWriter(fInfo)
	errWriter := zerolog.MultiLevelWriter(fError)

	errLoggerWriter := &LoggerWriter{w: errWriter, level: zerolog.ErrorLevel}
	infoLoggerWriter := &LoggerWriter{w: infoWriter, level: zerolog.InfoLevel}

	w := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, errLoggerWriter, infoLoggerWriter)
	logger := zerolog.New(w).With().Timestamp().Logger()

	return &logger
}
