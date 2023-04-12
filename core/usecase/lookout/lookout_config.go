package lookout

type LookoutConfig struct {
	Id          int
	Name        string
	Query       string
	Cron        string
	NotifyLocal bool
	NotifyMail  bool
}

type LookoutConfigCreate struct {
	Name        string
	Query       string
	Cron        string
	NotifyLocal bool
	NotifyMail  bool
}
