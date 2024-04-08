package stringx

import (
	"unsafe"
)

// 确保传入的字符串和字节切片的生命周期足够长，不会在转换后被释放或修改。
// 确保传入的字符串和字节切片的长度和容量是一致的，否则可能导致访问越界。
// 不要对转换后的字节切片或字符串进行修改，因为它们可能与原始的字符串或字节切片共享底层的内存。

// UnsafeToBytes 非安全的把 string 转换为 []byte 它必须遵守上述规则.
func UnsafeToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// UnsafeToString 非安全的把 []byte 转换为 string 它必须遵守上述规则.
func UnsafeToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
