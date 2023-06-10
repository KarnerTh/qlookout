package orchestration

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/KarnerTh/query-lookout/core/config"
)

func setupConfig() config.Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.WithError(err).Fatal("Could not get home dir")
	}

	viper.AddConfigPath(homeDir)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("QL")
	viper.AutomaticEnv()
	viper.SetConfigName(".query-lookout")

	err = viper.ReadInConfig()
	if err != nil {
		log.WithError(err).Fatal("Could not read config")
	}

	return config.New(homeDir)
}
