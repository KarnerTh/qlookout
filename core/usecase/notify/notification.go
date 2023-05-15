package notify

import (
	"time"

	"github.com/KarnerTh/query-lookout/observer"
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
