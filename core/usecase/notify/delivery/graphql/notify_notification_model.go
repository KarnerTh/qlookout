package graphql

import (
	"time"

	"github.com/KarnerTh/query-lookout/usecase/notify"
)

type notificationModel interface {
	Title() string
	Description() string
	DeepLink() string
	Timestamp() string
}

type notificationModelResolver struct {
	notification notify.Notification
}

func (r notificationModelResolver) Title() string {
	return r.notification.Title
}

func (r notificationModelResolver) Description() string {
	return r.notification.Description
}

func (r notificationModelResolver) DeepLink() string {
	return r.notification.DeepLink
}

func (r notificationModelResolver) Timestamp() string {
	return r.notification.Timestamp.UTC().Format(time.RFC3339)
}
