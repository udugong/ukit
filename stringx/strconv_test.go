package stringx

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestUnsafeToBytes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []byte
	}{
		{
			name: "normal",
			s:    "Hello, gopher!",
			want: []byte("Hello, gopher!"),
		},
		{
			name: "normal_emoji",
			s:    "😀Hello, gopher!",
			want: []byte("😀Hello, gopher!"),
		},
		{
			name: "normal_rune",
			s:    "你好！",
			want: []byte("你好！"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, UnsafeToBytes(tt.s))
		})
	}
}

func TestUnsafeToString(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want string
	}{
		{
			name: "normal",
			b:    []byte("Hello, gopher!"),
			want: "Hello, gopher!",
		},
		{
			name: "normal_emoji",
			b:    []byte("😀Hello, gopher!"),
			want: "😀Hello, gopher!",
		},
		{
			name: "normal_rune",
			b:    []byte("你好！"),
			want: "你好！",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, UnsafeToString(tt.b))
		})
	}
}

// goos: windows
// goarch: amd64
// pkg: github.com/udugong/ukit/stringx
// cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
// BenchmarkStringToBytes/unsafe-12      1000000000   0.2555 ns/op   0 B/op   0 allocs/op
// BenchmarkStringToBytes/unsafe_v1-12   1000000000   0.2555 ns/op   0 B/op   0 allocs/op
// BenchmarkStringToBytes/[]byte-12      1000000000   0.2550 ns/op   0 B/op   0 allocs/op
func BenchmarkStringToBytes(b *testing.B) {
	const str = "Hello, gopher! Hello, gopher! Hello, gopher!"

	b.Run("unsafe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = UnsafeToBytes(str)
		}
	})
	b.Run("unsafe_v1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = unsafeStingToBytesV1(str)
		}
	})
	b.Run("[]byte()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = []byte(str)
		}
	})
}

func unsafeStingToBytesV1(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// goos: windows
// goarch: amd64
// pkg: github.com/udugong/ukit/stringx
// cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
// BenchmarkBytesToString/unsafe-12      1000000000   0.2545 ns/op   0 B/op   0 allocs/op
// BenchmarkBytesToString/unsafe_v1-12   1000000000   0.2550 ns/op   0 B/op   0 allocs/op
// BenchmarkBytesToString/string()-12    50000415    22.71 ns/op    48 B/op   1 allocs/op
func BenchmarkBytesToString(b *testing.B) {
	var bytes = []byte("Hello, gopher! Hello, gopher! Hello, gopher!")

	b.Run("unsafe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = UnsafeToString(bytes)
		}
	})
	b.Run("unsafe_v1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = unsafeBytesToStringV1(bytes)
		}
	})
	b.Run("string()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = string(bytes)
		}
	})
}

func unsafeBytesToStringV1(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}
