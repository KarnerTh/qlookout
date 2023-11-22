package orchestration

import (
	"log/slog"
	"os"
	"strings"

	"github.com/KarnerTh/query-lookout/core/config"
)

func setupLogger(config config.Config) {
	logLevel := parseLevel(config.LogLevel())
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(logger)
}

func parseLevel(lvl string) slog.Level {
	switch strings.ToLower(lvl) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
