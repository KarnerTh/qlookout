package notify

type Notifier interface {
	Send(value Notification) error
}
