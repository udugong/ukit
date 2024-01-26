package option

type Option[T any] interface {
	Apply(*T)
}

type funcOption[T any] func(*T)

func (f funcOption[T]) Apply(t *T) {
	f(t)
}

func NewFuncOption[T any](f func(*T)) funcOption[T] {
	return f
}
