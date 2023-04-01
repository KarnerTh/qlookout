package config

import "github.com/spf13/viper"

const (
	dataSource      = "data_source"
	baseUrl         = "base_url"
	mailFromAddress = "mail_from_address"
	mailToAddress   = "mail_to_address"
	mailSmtpHost    = "mail_smtp_host"
	mailSmtpPort    = "mail_smtp_port"
)

type config struct{}

type Config interface {
	DataSource() string
	BaseUrl() string
	MailFromAddress() string
	MailToAddress() string
	MailSmtpHost() string
	MailSmtpPort() string
}

func New() Config {
	return config{}
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