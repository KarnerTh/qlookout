package orchestration

import (
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/config"
)

func setupLogger(config config.Config) {
	logLevel, err := log.ParseLevel(config.LogLevel())
	if err != nil {
		log.WithError(err).Errorf("Could not parse log level %s, fallback to INFO", config.LogLevel())
		logLevel = log.WarnLevel
	}

	log.SetLevel(logLevel)
}
