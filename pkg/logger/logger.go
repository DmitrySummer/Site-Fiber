package logger

import (
	"dl/new-web-site/config"
	"log/slog"
	"os"
)

func NewLogger(conf *config.LogConfig) *slog.Logger {
	var handler slog.Handler

	if conf.Format == "json" {
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.Level(conf.Level),
		})
	} else {
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.Level(conf.Level),
		})
	}

	logger := slog.New(handler)
	return logger
}
