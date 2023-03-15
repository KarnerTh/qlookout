package orchestration

import log "github.com/sirupsen/logrus"

func setupLogger() {
	log.SetLevel(log.DebugLevel)
}
