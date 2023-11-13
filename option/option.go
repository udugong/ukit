package option

type Option[T any] interface {
	Apply(*T)
}

type funcOption[T any] struct {
	f func(*T)
}

func (fdo *funcOption[T]) Apply(do *T) {
	fdo.f(do)
}

func NewFuncOption[T any](f func(*T)) *funcOption[T] {
	return &funcOption[T]{
		f: f,
	}
}
