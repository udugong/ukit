package reverse

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	type testCase[T any] struct {
		name string
		src  []T
		want []T
	}
	tests := []testCase[rune]{
		{
			name: "normal",
			src:  []rune("你好123"),
			want: []rune("321好你"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice(tt.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
