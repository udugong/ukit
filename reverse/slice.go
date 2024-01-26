package reverse

// Slice 反转切片中的元素.
func Slice[T any](src []T) []T {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
	return src
}
