package graphql

type lookoutConfigModel struct {
	Id          int32
	Name        string
	Query       string
	Cron        string
	NotifyLocal bool
	NotifyMail  bool
}

type lookoutConfigCreateModel struct {
	Name        string
	Query       string
	Cron        string
	NotifyLocal bool
	NotifyMail  bool
}
