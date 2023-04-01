package config

import "github.com/spf13/viper"

const (
	dataSource = "data_source"
)

type config struct{}

type Config interface {
	DataSource() string
}

func New() Config {
	return config{}
}

func (c config) DataSource() string {
	return viper.GetString(dataSource)
}
