package watch

type cronJob[T any] struct {
	value   T
	execute func(value T)
}

func (c cronJob[T]) Run() {
	c.execute(c.value)
}
