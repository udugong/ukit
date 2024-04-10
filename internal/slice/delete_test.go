package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/udugong/ukit/internal/errs"
)

func TestDeleteByIter(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name    string
		src     S
		index   int
		want    S
		wantErr error
	}
	tests := []testCase[[]int, int]{
		{
			name:  "delete_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 2,
			want:  []int{0, 1, 3, 4},
		},
		{
			name:  "delete_last_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 4,
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "delete_first_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 0,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:    "delete_not_exists_idx",
			src:     []int{0, 1, 2, 3, 4},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:    "out_of_range",
			src:     []int{},
			index:   0,
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteByIter(tt.src, tt.index)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeleteByAppend(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name    string
		src     S
		index   int
		want    S
		wantErr error
	}
	tests := []testCase[[]int, int]{
		{
			name:  "delete_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 2,
			want:  []int{0, 1, 3, 4},
		},
		{
			name:  "delete_last_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 4,
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "delete_first_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 0,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:    "delete_not_exists_idx",
			src:     []int{0, 1, 2, 3, 4},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:    "out_of_range",
			src:     []int{},
			index:   0,
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteByAppend(tt.src, tt.index)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeleteByCopy(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name    string
		src     S
		index   int
		want    S
		wantErr error
	}
	tests := []testCase[[]int, int]{
		{
			name:  "delete_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 2,
			want:  []int{0, 1, 3, 4},
		},
		{
			name:  "delete_last_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 4,
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "delete_first_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: 0,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:    "delete_not_exists_idx",
			src:     []int{0, 1, 2, 3, 4},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:    "out_of_range",
			src:     []int{},
			index:   0,
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteByCopy(tt.src, tt.index)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

// goos: windows
// goarch: amd64
// pkg: github.com/udugong/ukit/internal/slice
// cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
// BenchmarkDelete/delete_by_iter_8-12       36922735    30.86 ns/op    64 B/op   1 allocs/op
// BenchmarkDelete/delete_by_append_8-12     37498358    31.20 ns/op    64 B/op   1 allocs/op
// BenchmarkDelete/delete_by_copy_8-12       36364627    32.17 ns/op    64 B/op   1 allocs/op
// BenchmarkDelete/delete_by_iter_16-12      27907495    43.86 ns/op   128 B/op   1 allocs/op
// BenchmarkDelete/delete_by_append_16-12    26965928    40.87 ns/op   128 B/op   1 allocs/op
// BenchmarkDelete/delete_by_copy_16-12      28234230    40.78 ns/op   128 B/op   1 allocs/op
// BenchmarkDelete/delete_by_iter_32-12      18181762    63.69 ns/op   256 B/op   1 allocs/op
// BenchmarkDelete/delete_by_append_32-12    18604189    60.52 ns/op   256 B/op   1 allocs/op
// BenchmarkDelete/delete_by_copy_32-12      19561795    61.19 ns/op   256 B/op   1 allocs/op
// BenchmarkDelete/delete_by_iter_64-12      10810780   111.1 ns/op    512 B/op   1 allocs/op
// BenchmarkDelete/delete_by_append_64-12    11594235   104.9 ns/op    512 B/op   1 allocs/op
// BenchmarkDelete/delete_by_copy_64-12      11215068   107.4 ns/op    512 B/op   1 allocs/op
// BenchmarkDelete/delete_by_iter_128-12      5429871   213.1 ns/op   1024 B/op   1 allocs/op
// BenchmarkDelete/delete_by_append_128-12    6122482   193.7 ns/op   1024 B/op   1 allocs/op
// BenchmarkDelete/delete_by_copy_128-12      5940590   194.7 ns/op   1024 B/op   1 allocs/op
// BenchmarkDelete/delete_by_iter_512-12      1523808   796.3 ns/op   4096 B/op   1 allocs/op
// BenchmarkDelete/delete_by_append_512-12    1688952   721.2 ns/op   4096 B/op   1 allocs/op
// BenchmarkDelete/delete_by_copy_512-12      1638223   736.2 ns/op   4096 B/op   1 allocs/op
func BenchmarkDelete(b *testing.B) {
	const (
		srcLen8   = 1 << 3
		srcLen16  = 1 << 4
		srcLen32  = 1 << 5
		srcLen64  = 1 << 6
		srcLen128 = 1 << 7
		srcLen512 = 1 << 9
	)
	getSlice := func(n int) []int {
		return make([]int, n)
	}

	b.Run("delete_by_iter_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen8)
			_, _ = DeleteByIter(src, srcLen8/2)
		}
	})
	b.Run("delete_by_append_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen8)
			_, _ = DeleteByAppend(src, srcLen8/2)
		}
	})
	b.Run("delete_by_copy_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen8)
			_, _ = DeleteByCopy(src, srcLen8/2)
		}
	})

	b.Run("delete_by_iter_16", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen16)
			_, _ = DeleteByIter(src, srcLen16/2)
		}
	})
	b.Run("delete_by_append_16", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen16)
			_, _ = DeleteByAppend(src, srcLen16/2)
		}
	})
	b.Run("delete_by_copy_16", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen16)
			_, _ = DeleteByCopy(src, srcLen16/2)
		}
	})

	b.Run("delete_by_iter_32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen32)
			_, _ = DeleteByIter(src, srcLen32/2)
		}
	})
	b.Run("delete_by_append_32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen32)
			_, _ = DeleteByAppend(src, srcLen32/2)
		}
	})
	b.Run("delete_by_copy_32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen32)
			_, _ = DeleteByCopy(src, srcLen32/2)
		}
	})

	b.Run("delete_by_iter_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen64)
			_, _ = DeleteByIter(src, srcLen64/2)
		}
	})
	b.Run("delete_by_append_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen64)
			_, _ = DeleteByAppend(src, srcLen64/2)
		}
	})
	b.Run("delete_by_copy_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen64)
			_, _ = DeleteByCopy(src, srcLen64/2)
		}
	})

	b.Run("delete_by_iter_128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen128)
			_, _ = DeleteByIter(src, srcLen128/2)
		}
	})
	b.Run("delete_by_append_128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen128)
			_, _ = DeleteByAppend(src, srcLen128/2)
		}
	})
	b.Run("delete_by_copy_128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen128)
			_, _ = DeleteByCopy(src, srcLen128/2)
		}
	})

	b.Run("delete_by_iter_512", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen512)
			_, _ = DeleteByIter(src, srcLen512/2)
		}
	})
	b.Run("delete_by_append_512", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen512)
			_, _ = DeleteByAppend(src, srcLen512/2)
		}
	})
	b.Run("delete_by_copy_512", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			src := getSlice(srcLen512)
			_, _ = DeleteByCopy(src, srcLen512/2)
		}
	})
}
