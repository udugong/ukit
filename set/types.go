package set

type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	// Exists 返回是否存在这个元素
	Exists(key T) bool
	Keys() []T
}
