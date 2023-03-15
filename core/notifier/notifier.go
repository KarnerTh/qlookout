package notifier

import "sync"

// TODO: add unsubscribe?
type Notifier[T any] interface {
	Publish(value T)
	Subscribe() <-chan T
}

type notifier[T any] struct {
	mu   sync.Mutex
	subs []chan T
}

func New[T any]() Notifier[T] {
	return &notifier[T]{
		subs: []chan T{},
	}
}

func (n *notifier[T]) Publish(value T) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, ch := range n.subs {
		ch <- value
	}
}

func (n *notifier[T]) Subscribe() <-chan T {
	n.mu.Lock()
	defer n.mu.Unlock()

	ch := make(chan T)
	n.subs = append(n.subs, ch)
	return ch
}
