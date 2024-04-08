package stringx

import "github.com/udugong/ukit/slice"

func Reverse(s string) string {
	bytes := []byte(s)
	slice.ReverseSelf(bytes)
	return UnsafeToString(bytes)
}

func RuneReverse(s string) string {
	runes := []rune(s)
	slice.ReverseSelf(runes)
	return string(runes)
}
