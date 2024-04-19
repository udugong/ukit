package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/udugong/ukit/internal/errs"
)

func TestDelete(t *testing.T) {
	// Delete 主要依赖于 internal/slice.DeleteByAppend 来保证正确性
	type testCase[S ~[]E, E any] struct {
		name    string
		src     S
		index   int
		want    S
		wantErr error
	}
	tests := []testCase[[]int, int]{
		{
			name:    "normal",
			src:     []int{1, 2, 3, 4, 5},
			index:   0,
			want:    []int{2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name:    "out_of_range",
			src:     []int{1, 2, 3, 4, 5},
			index:   -1,
			want:    nil,
			wantErr: errs.NewErrIndexOutOfRange(5, -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete(tt.src, tt.index)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
