package graphql

type lookoutConfigUpdateModel struct {
	Name        *string
	Query       *string
	Cron        *string
	NotifyLocal *bool
	NotifyMail  *bool
}
