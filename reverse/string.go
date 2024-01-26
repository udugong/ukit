package reverse

// String 反转 ASCII 字符串.
func String(s string) string {
	return string(Slice([]byte(s)))
}

// RuneString 反转 Unicode 字符串.
func RuneString(s string) string {
	return string(Slice([]rune(s)))
}
