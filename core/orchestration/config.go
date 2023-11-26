package orchestration

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"

	"github.com/KarnerTh/query-lookout/core/config"
)

func setupConfig(configPath string) config.Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Could not get home dir", slog.Any("error", err))
		panic("Could not get home dir")
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

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not required
			slog.Info("Starting without config file")
		} else {
			slog.Error("Could not read config", slog.Any("error", err))
			panic("Could not read config")
		}
	}

	return config.New(homeDir)
}
