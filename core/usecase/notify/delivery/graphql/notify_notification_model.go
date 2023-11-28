package graphql

import (
	"time"

	"github.com/KarnerTh/qlookout/core/usecase/notify"
)

type notificationModel interface {
	LookoutId() int32
	RuleId() int32
	Title() string
	Description() string
	Timestamp() string
}

type notificationModelResolver struct {
	notification notify.Notification
}

func (r notificationModelResolver) LookoutId() int32 {
	return int32(r.notification.LookoutId)
}

func (r notificationModelResolver) RuleId() int32 {
	return int32(r.notification.RuleId)
}

func (r notificationModelResolver) Title() string {
	return r.notification.Title
}

func (r notificationModelResolver) Description() string {
	return r.notification.Description
}

func (r notificationModelResolver) Timestamp() string {
	return r.notification.Timestamp.UTC().Format(time.RFC3339)
}
