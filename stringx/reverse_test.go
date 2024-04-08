package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal",
			s:    "hello world",
			want: "dlrow olleh",
		},
		{
			name: "normal_empty",
			s:    "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Reverse(tt.s))
		})
	}
}

func TestRuneReverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal",
			s:    "你好",
			want: "好你",
		},
		{
			name: "normal_empty",
			s:    "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RuneReverse(tt.s))
		})
	}
}

// goos: windows
// goarch: amd64
// pkg: github.com/udugong/ukit/stringx
// cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
// BenchmarkReverse-12   156148360   7.707 ns/op   0 B/op   0 allocs/op
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Reverse("hello world")
	}
}

// goos: windows
// goarch: amd64
// pkg: github.com/udugong/ukit/stringx
// cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
// BenchmarkRuneReverse-12   14723943   81.50 ns/op   0 B/op   0 allocs/op
func BenchmarkRuneReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RuneReverse("hello world")
	}
}
