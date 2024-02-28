package heap

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func (h Heap[T]) verify(t *testing.T, i int) {
	t.Helper()
	n := h.Len()
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if h.Less(j1, i) {
			t.Errorf("heap invariant invalidated [%v] = %v > [%v] = %v", i, h[i], j1, h[j1])
			return
		}
		h.verify(t, j1)
	}
	if j2 < n {
		if h.Less(j2, i) {
			t.Errorf("heap invariant invalidated [%v] = %v > [%v] = %v", i, h[i], j1, h[j2])
			return
		}
		h.verify(t, j2)
	}
}

func TestInit0(t *testing.T) {
	h := new(Heap[int])
	for i := 20; i > 0; i-- {
		h.Push(0) // all elements are the same
	}
	h.Init()
	h.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x := h.Pop()
		h.verify(t, 0)
		if x != 0 {
			t.Errorf("%v.th pop got %v; want %v", i, x, 0)
		}
	}
}

func TestInit1(t *testing.T) {
	h := new(Heap[int])
	for i := 20; i > 0; i-- {
		h.Push(i) // all elements are different
	}
	h.Init()
	h.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x := h.PopElement()
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func Test(t *testing.T) {
	h := new(Heap[int])
	h.verify(t, 0)

	for i := 20; i > 10; i-- {
		h.Push(i)
	}
	h.Init()
	h.verify(t, 0)

	for i := 10; i > 0; i-- {
		h.PushElement(i)
		h.verify(t, 0)
	}

	for i := 1; h.Len() > 0; i++ {
		x := h.PopElement()
		if i < 20 {
			h.PushElement(20 + i)
		}
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestRemove0(t *testing.T) {
	h := new(Heap[int])
	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	h.verify(t, 0)

	for h.Len() > 0 {
		i := h.Len() - 1
		x := h.Remove(i)
		if x != i {
			t.Errorf("Remove(%d) got %d; want %d", i, x, i)
		}
		h.verify(t, 0)
	}
}

func TestRemove1(t *testing.T) {
	h := new(Heap[int])
	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	h.verify(t, 0)

	for i := 0; h.Len() > 0; i++ {
		x := h.Remove(0)
		if x != i {
			t.Errorf("Remove(0) got %d; want %d", x, i)
		}
		h.verify(t, 0)
	}
}

func TestRemove2(t *testing.T) {
	N := 10

	h := new(Heap[int])
	for i := 0; i < N; i++ {
		h.Push(i)
	}
	h.verify(t, 0)

	m := make(map[int]bool)
	for h.Len() > 0 {
		m[h.Remove((h.Len()-1)/2)] = true
		h.verify(t, 0)
	}

	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func BenchmarkDup(b *testing.B) {
	const n = 10000
	h := make(Heap[int], 0, n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			h.PushElement(0) // all elements are the same
		}
		for h.Len() > 0 {
			h.PopElement()
		}
	}
}

func TestFix(t *testing.T) {
	h := new(Heap[int])
	h.verify(t, 0)

	for i := 200; i > 0; i -= 10 {
		h.PushElement(i)
	}
	h.verify(t, 0)

	if (*h)[0] != 10 {
		t.Fatalf("Expected head to be 10, was %d", (*h)[0])
	}
	(*h)[0] = 210
	h.Fix(0)
	h.verify(t, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(h.Len())
		if i&1 == 0 {
			(*h)[elem] *= 2
		} else {
			(*h)[elem] /= 2
		}
		h.Fix(elem)
		h.verify(t, 0)
	}
}

func TestNewHeap(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name     string
		capacity int
		val      []T
		want     *Heap[T]
	}
	h := make(Heap[int], 0, 3)
	tests := []testCase[int]{
		{
			name:     "normal",
			capacity: 3,
			val:      []int{},
			want:     &h,
		},
		{
			name:     "has_val",
			capacity: 0,
			val:      []int{2, 1, 3},
			want:     &Heap[int]{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewHeap(tt.capacity, tt.val...))
		})
	}
}

func TestHeap_Replace(t *testing.T) {
	h := new(Heap[int])
	h.verify(t, 0)

	for i := 200; i > 0; i -= 10 {
		h.PushElement(i)
	}
	h.verify(t, 0)

	if (*h)[0] != 10 {
		t.Fatalf("Expected head to be 10, was %d", (*h)[0])
	}

	h.Replace(0, 210)
	h.verify(t, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(h.Len())
		if i&1 == 0 {
			h.Replace(elem, (*h)[elem]*2)
		} else {
			h.Replace(elem, (*h)[elem]/2)
		}
		h.verify(t, 0)
	}
}
