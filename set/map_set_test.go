package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSet_Add(t *testing.T) {
	addVales := []int{1, 2, 3, 1}
	s := make(MapSet[int], 3)
	t.Run("Add", func(t *testing.T) {
		for _, val := range addVales {
			s.Add(val)
		}
		assert.Equal(t, s, MapSet[int]{
			1: {},
			2: {},
			3: {},
		})
	})
}

func TestMapSet_Delete(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		set     MapSet[T]
		key     T
		wantSet MapSet[T]
	}
	tests := []testCase[int]{
		{
			name:    "normal",
			set:     MapSet[int]{1: {}},
			key:     1,
			wantSet: MapSet[int]{},
		},
		{
			name:    "normal",
			set:     MapSet[int]{1: {}},
			key:     2,
			wantSet: MapSet[int]{1: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Delete(tt.key)
			assert.Equal(t, tt.wantSet, tt.set)
		})
	}
}

func TestMapSet_Exists(t *testing.T) {
	s := make(MapSet[int], 1)
	s.Add(1)
	type testCase[T comparable] struct {
		name string
		key  T
		want bool
	}
	tests := []testCase[int]{
		{
			name: "is_exists",
			key:  1,
			want: true,
		},
		{
			name: "not_exists",
			key:  2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.Exists(tt.key))
		})
	}
}

func TestMapSet_Keys(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		set  MapSet[T]
	}
	tests := []testCase[int]{
		{
			name: "normal",
			set: MapSet[int]{
				1: {},
				2: {},
				3: {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, true, equal(tt.set.Keys(), tt.set))
		})
	}
}

func equal(nums []int, m map[int]struct{}) bool {
	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			return false
		}
		delete(m, num)
	}
	return len(m) == 0
}

func BenchmarkMapSet(b *testing.B) {
	const n = 100
	s := make(MapSet[int], 100)
	b.Run("map_set_add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			do(n, func(i int) {
				s.Add(i)
			})
		}
	})
	b.Run("map_set_exists", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			do(n, func(i int) {
				s.Exists(i)
			})
		}
	})
	b.Run("map_set_delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			do(n, func(i int) {
				s.Delete(i)
			})
		}
	})
}

func do(n int, op func(i int)) {
	for i := 0; i < n; i++ {
		op(i)
	}
}
