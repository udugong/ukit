package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvToAny(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		name string
		src  S
		want []any
	}
	tests := []testCase[[]string, string]{
		{
			name: "normal",
			src:  []string{"a", "b", "c"},
			want: []any{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ConvToAny(tt.src))
		})
	}
}
