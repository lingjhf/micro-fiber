package microfiber

type Database[T any] interface {
	Connect() (T, error)
}

func New[R any](config Database[R]) Database[R] {
	return config
}
