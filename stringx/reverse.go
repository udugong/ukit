package stringx

import "github.com/udugong/ukit/slicex"

func Reverse(s string) string {
	bytes := []byte(s)
	slicex.ReverseSelf(bytes)
	return UnsafeToString(bytes)
}

func RuneReverse(s string) string {
	runes := []rune(s)
	slicex.ReverseSelf(runes)
	return string(runes)
}
