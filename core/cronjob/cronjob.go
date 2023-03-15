package cronjob

type CronJob[T any] struct {
	Value   T
	Execute func(value T)
}

func (c CronJob[T]) Run() {
	c.Execute(c.Value)
}
