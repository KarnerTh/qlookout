package notify

//go:generate mockery --name Notifier 
type Notifier interface {
	Send(value Notification) error
}
