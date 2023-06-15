package orchestration

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/KarnerTh/query-lookout/core/config"
)

func setupConfig(configPath string) config.Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.WithError(err).Fatal("Could not get home dir")
	}

	viper.SetConfigType("yaml")

	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.AddConfigPath(homeDir)
		viper.SetEnvPrefix("QL")
		viper.AutomaticEnv()
		viper.SetConfigName(".query-lookout")
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.WithError(err).Fatal("Could not read config")
	}

	return config.New(homeDir)
}
