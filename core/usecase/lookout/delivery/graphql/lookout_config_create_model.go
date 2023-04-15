package graphql

type lookoutConfigCreateModel struct {
	Name        string
	Query       string
	Cron        string
	NotifyLocal bool
	NotifyMail  bool
}
