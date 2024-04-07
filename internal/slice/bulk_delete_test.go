package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBulkDeleteByIter(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name  string
		src   S
		index []int
		want  S
	}
	tests := []testCase[[]int, int]{
		{
			name:  "delete_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{3, 1, 4},
			want:  []int{0, 2},
		},
		{
			name:  "delete_empty_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{},
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			name:  "delete_first_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{0},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "delete_last_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{4},
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "delete_repeat_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{1, 3, 4, 1},
			want:  []int{0, 2},
		},
		{
			name:  "delete_not_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{5, -1},
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			name:  "delete_not_exists_idx_&_out_of_range",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{1, 3, 4, 1, 9, 10, 11, 12},
			want:  []int{0, 2},
		},
		{
			name:  "delete_empty_slice",
			src:   []int{},
			index: []int{1},
			want:  []int{},
		},
		{
			name:  "delete_empty_slice_&_empty_idx",
			src:   []int{},
			index: []int{},
			want:  []int{},
		},
		{
			name:  "delete_all_element",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{0, 1, 2, 3, 4},
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, BulkDeleteByIter(tt.src, tt.index...))
		})
	}
}

func TestBulkDeleteByAppend(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name  string
		src   S
		index []int
		want  S
	}
	tests := []testCase[[]int, int]{
		{
			name:  "delete_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{3, 1, 4},
			want:  []int{0, 2},
		},
		{
			name:  "delete_empty_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{},
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			name:  "delete_first_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{0},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "delete_last_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{4},
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "delete_repeat_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{1, 3, 4, 1},
			want:  []int{0, 2},
		},
		{
			name:  "delete_not_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{5, -1},
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			name:  "delete_not_exists_idx_&_out_of_range",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{1, 3, 4, 1, 9, 10, 11, 12},
			want:  []int{0, 2},
		},
		{
			name:  "delete_empty_slice",
			src:   []int{},
			index: []int{1},
			want:  []int{},
		},
		{
			name:  "delete_empty_slice_&_empty_idx",
			src:   []int{},
			index: []int{},
			want:  []int{},
		},
		{
			name:  "delete_all_element",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{0, 1, 2, 3, 4},
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, BulkDeleteByAppend(tt.src, tt.index...))
		})
	}
}

func TestBulkDeleteByCopy(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name  string
		src   S
		index []int
		want  S
	}
	tests := []testCase[[]int, int]{
		{
			name:  "delete_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{3, 1, 4},
			want:  []int{0, 2},
		},
		{
			name:  "delete_empty_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{},
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			name:  "delete_first_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{0},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "delete_last_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{4},
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "delete_repeat_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{1, 3, 4, 1},
			want:  []int{0, 2},
		},
		{
			name:  "delete_not_exists_idx",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{5, -1},
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			name:  "delete_not_exists_idx_&_out_of_range",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{1, 3, 4, 1, 9, 10, 11, 12},
			want:  []int{0, 2},
		},
		{
			name:  "delete_empty_slice",
			src:   []int{},
			index: []int{1},
			want:  []int{},
		},
		{
			name:  "delete_empty_slice_&_empty_idx",
			src:   []int{},
			index: []int{},
			want:  []int{},
		},
		{
			name:  "delete_all_element",
			src:   []int{0, 1, 2, 3, 4},
			index: []int{0, 1, 2, 3, 4},
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, BulkDeleteByCopy(tt.src, tt.index...))
		})
	}
}

func getNeedDeleteIdx(length int) []int {
	res := make([]int, 0, length)
	for i := 0; i < length; i++ {
		res = append(res, length-i)
	}
	return res
}

func BenchmarkBulkDelete(b *testing.B) {
	const (
		srcLen8  = 1 << 3
		srcLen64 = 1 << 6
		delIdx2  = 1 << 1
		delIdx4  = 1 << 2
		delIdx8  = 1 << 3
		delIdx16 = 1 << 4
		delIdx32 = 1 << 5
	)
	getSlice := func(n int) []int {
		return make([]int, n)
	}

	b.Run("bulk_delete_2_by_iter_length_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen8)
			idx := getNeedDeleteIdx(delIdx2)
			_ = BulkDeleteByIter(s, idx...)
		}
	})
	b.Run("bulk_delete_2_by_append_length_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen8)
			idx := getNeedDeleteIdx(delIdx2)
			_ = BulkDeleteByAppend(s, idx...)
		}
	})
	b.Run("bulk_delete_2_by_copy_length_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen8)
			idx := getNeedDeleteIdx(delIdx2)
			_ = BulkDeleteByCopy(s, idx...)
		}
	})

	b.Run("bulk_delete_4_by_iter_length_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen8)
			idx := getNeedDeleteIdx(delIdx4)
			_ = BulkDeleteByIter(s, idx...)
		}
	})
	b.Run("bulk_delete_4_by_append_length_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen8)
			idx := getNeedDeleteIdx(delIdx4)
			_ = BulkDeleteByAppend(s, idx...)
		}
	})
	b.Run("bulk_delete_4_by_copy_length_8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen8)
			idx := getNeedDeleteIdx(delIdx4)
			_ = BulkDeleteByCopy(s, idx...)
		}
	})

	b.Run("bulk_delete_2_by_iter_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx2)
			_ = BulkDeleteByIter(s, idx...)
		}
	})
	b.Run("bulk_delete_2_by_append_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx2)
			_ = BulkDeleteByAppend(s, idx...)
		}
	})
	b.Run("bulk_delete_2_by_copy_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx2)
			_ = BulkDeleteByCopy(s, idx...)
		}
	})

	b.Run("bulk_delete_4_by_iter_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx4)
			_ = BulkDeleteByIter(s, idx...)
		}
	})
	b.Run("bulk_delete_4_by_append_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx4)
			_ = BulkDeleteByAppend(s, idx...)
		}
	})
	b.Run("bulk_delete_4_by_copy_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx4)
			_ = BulkDeleteByCopy(s, idx...)
		}
	})

	b.Run("bulk_delete_8_by_iter_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx8)
			_ = BulkDeleteByIter(s, idx...)
		}
	})
	b.Run("bulk_delete_8_by_append_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx8)
			_ = BulkDeleteByAppend(s, idx...)
		}
	})
	b.Run("bulk_delete_8_by_copy_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx8)
			_ = BulkDeleteByCopy(s, idx...)
		}
	})

	b.Run("bulk_delete_16_by_iter_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx16)
			_ = BulkDeleteByIter(s, idx...)
		}
	})
	b.Run("bulk_delete_16_by_append_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx16)
			_ = BulkDeleteByAppend(s, idx...)
		}
	})
	b.Run("bulk_delete_16_by_copy_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx16)
			_ = BulkDeleteByCopy(s, idx...)
		}
	})

	b.Run("bulk_delete_32_by_iter_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx32)
			_ = BulkDeleteByIter(s, idx...)
		}
	})
	b.Run("bulk_delete_32_by_append_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx32)
			_ = BulkDeleteByAppend(s, idx...)
		}
	})
	b.Run("bulk_delete_32_by_copy_length_64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := getSlice(srcLen64)
			idx := getNeedDeleteIdx(delIdx32)
			_ = BulkDeleteByCopy(s, idx...)
		}
	})
}
