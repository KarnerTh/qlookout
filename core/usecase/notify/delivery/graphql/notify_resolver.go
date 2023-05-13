package graphql

import (
	"context"
	"time"

	"github.com/KarnerTh/query-lookout/usecase/notify"
)

type NotifyResolver struct {
	events     chan notificationModel
	subscriber chan *subscriber
}

// Source: https://github.com/matiasanaya/go-graphql-subscription-example
type subscriber struct {
	stop   <-chan struct{}
	events chan<- notificationModel
}

func NewNotifyResolver(notificationSubscriber notify.NotificationSubscriber) NotifyResolver {
	r := NotifyResolver{
		events:     make(chan notificationModel),
		subscriber: make(chan *subscriber),
	}

	go r.broadcastNotifications()
	go r.setupNotificationSubscriber(notificationSubscriber)

	return r
}

func (r *NotifyResolver) setupNotificationSubscriber(notificationSubscriber notify.NotificationSubscriber) {
	c := notificationSubscriber.Subscribe()
	for notification := range c {
		r.events <- notificationModelResolver{notification: notification}
	}
}

func (r *NotifyResolver) broadcastNotifications() {
	subscribers := map[int64]*subscriber{}
	unsubscribe := make(chan int64)

	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.subscriber:
			subscribers[time.Now().UnixMilli()] = s
		case event := <-r.events:
			for id, s := range subscribers {
				go func(id int64, s *subscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- event:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}

func (r *NotifyResolver) NewNotification(ctx context.Context) <-chan notificationModel {
	c := make(chan notificationModel)
	r.subscriber <- &subscriber{events: c, stop: ctx.Done()}
	return c
}
