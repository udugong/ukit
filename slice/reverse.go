package slice

// Reverse 创建一个新的 slice 存储反转后 src 的元素,并不会修改 src.
func Reverse[S ~[]E, E any](src S) S {
	dst := make(S, 0, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		dst = append(dst, src[i])
	}
	return dst
}

// ReverseSelf 直接在 src 上反转切片中的元素.
func ReverseSelf[S ~[]E, E any](src S) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}
