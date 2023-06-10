package config

import (
	"path"

	"github.com/spf13/viper"
)

const (
	logLevel        = "log_level"
	databaseFile    = "database_file"
	dataSource      = "data_source"
	baseUrl         = "base_url"
	mailFromAddress = "mail_from_address"
	mailToAddress   = "mail_to_address"
	mailSmtpHost    = "mail_smtp_host"
	mailSmtpPort    = "mail_smtp_port"
	mailUsername    = "mail_username"
	mailPassword    = "mail_password"
)

type config struct{}

//go:generate mockery --name Config
type Config interface {
	LogLevel() string
	DatabaseFile() string
	DataSource() string
	BaseUrl() string
	MailFromAddress() string
	MailToAddress() string
	MailSmtpHost() string
	MailSmtpPort() string
	MailUsername() string
	MailPassword() string
}

func New(homeDir string) Config {
	viper.SetDefault(logLevel, "INFO")
	viper.SetDefault(databaseFile, path.Join(homeDir, "data.db"))

	return config{}
}

func (c config) LogLevel() string {
	return viper.GetString(logLevel)
}

func (c config) DatabaseFile() string {
	return viper.GetString(databaseFile)
}

func (c config) DataSource() string {
	return viper.GetString(dataSource)
}

func (c config) BaseUrl() string {
	return viper.GetString(baseUrl)
}

func (c config) MailFromAddress() string {
	return viper.GetString(mailFromAddress)
}

func (c config) MailToAddress() string {
	return viper.GetString(mailToAddress)
}

func (c config) MailSmtpHost() string {
	return viper.GetString(mailSmtpHost)
}

func (c config) MailSmtpPort() string {
	return viper.GetString(mailSmtpPort)
}

func (c config) MailUsername() string {
	return viper.GetString(mailUsername)
}

func (c config) MailPassword() string {
	return viper.GetString(mailPassword)
}
