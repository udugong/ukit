package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name string
		src  S
		want S
	}
	tests := []testCase[[]int, int]{
		{
			name: "normal",
			src:  []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "empty",
			src:  []int{},
			want: []int{},
		},
		{
			name: "nil",
			src:  []int(nil),
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Reverse(tt.src))
		})
	}
}

func TestReverseSelf(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name string
		src  S
		want S
	}
	tests := []testCase[[]int, int]{
		{
			name: "normal",
			src:  []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "empty",
			src:  []int{},
			want: []int{},
		},
		{
			name: "nil",
			src:  []int(nil),
			want: []int(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSelf(tt.src)
			assert.Equal(t, tt.want, tt.src)
		})
	}
}

// goos: windows
// goarch: amd64
// pkg: github.com/udugong/ukit/slice
// cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
// BenchmarkReverse/reverse-12        43638267   26.20 ns/op   48 B/op   1 allocs/op
// BenchmarkReverse/reverse_self-12   631578614   1.883 ns/op   0 B/op   0 allocs/op
func BenchmarkReverse(b *testing.B) {
	src := []int{1, 2, 3, 4, 5}
	b.Run("reverse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Reverse(src)
		}
	})
	b.Run("reverse_self", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ReverseSelf(src)
		}
	})
}
