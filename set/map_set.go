package set

type MapSet[T comparable] map[T]struct{}

func New[T comparable](cap int) MapSet[T] {
	return make(MapSet[T], cap)
}

func (s MapSet[T]) Add(key T) {
	s[key] = struct{}{}
}

func (s MapSet[T]) Delete(key T) {
	delete(s, key)
}

func (s MapSet[T]) Exists(key T) bool {
	_, ok := s[key]
	return ok
}

func (s MapSet[T]) Keys() []T {
	ans := make([]T, 0, len(s))
	for key := range s {
		ans = append(ans, key)
	}
	return ans
}
