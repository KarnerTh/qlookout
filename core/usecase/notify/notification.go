package notify

import (
	"time"

	"github.com/KarnerTh/query-lookout/observer"
)

type NotificationPublisher = observer.Publisher[Notification]
type NotificationSubscriber = observer.Subscriber[Notification]

type Notification struct {
	Title       string
	Description string
	DeepLink    string
	Timestamp   time.Time
}
