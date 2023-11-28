package notify

import (
	"time"

	"github.com/KarnerTh/qlookout/core/observer"
)

type NotificationPublisher = observer.Publisher[Notification]
type NotificationSubscriber = observer.Subscriber[Notification]

type Notification struct {
	LookoutId   int
	RuleId      int
	Title       string
	Description string
	Timestamp   time.Time
}
