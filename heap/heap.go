// Package heap provides heap operations for any type that implements
// heap.Interface. A heap is a tree with the property that each node is the
// minimum-valued node in its subtree.
//
// The minimum element in the tree is the root, at index 0.
//
// A heap is a common way to implement a priority queue. To build a priority
// queue, implement the Heap interface with the (negative) priority as the
// ordering for the Less method, so Push adds items while Pop removes the
// highest-priority item from the queue.
package heap

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// Init has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that Push and Pop in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use heap.Push and heap.Pop.
type Interface[T any] interface {
	sort.Interface
	Push(x T) // add x as element Len()
	Pop() T   // remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init[T any](h Interface[T]) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push[T any](h Interface[T], x T) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop[T any](h Interface[T]) T {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove[T any](h Interface[T], i int) T {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix[T any](h Interface[T], i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up[T any](h Interface[T], j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down[T any](h Interface[T], i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

// Heap 小顶堆的实现.
// 大顶堆只需要把 Less() 方法改为 return (*h)[i] > (*h)[j] 即可.
type Heap[T constraints.Ordered] []T

// NewHeap 初始化一个小顶堆.
func NewHeap[T constraints.Ordered](capacity int, val ...T) *Heap[T] {
	length := len(val)
	if length > capacity {
		capacity = length
	}
	h := make(Heap[T], 0, capacity)
	if length > 0 {
		h = append(h, val...)
		Init[T](&h)
	}
	return &h
}

func (h *Heap[T]) Len() int           { return len(*h) }
func (h *Heap[T]) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *Heap[T]) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Push add x as element Len().
func (h *Heap[T]) Push(x T) {
	*h = append(*h, x)
}

// Pop remove and return element Len() - 1.
func (h *Heap[T]) Pop() (x T) {
	*h, x = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *Heap[T]) Init() {
	Init[T](h)
}

// PushElement 将元素x插入到堆中并进行上滤操作.
func (h *Heap[T]) PushElement(v T) {
	Push[T](h, v)
}

// PopElement 从堆中移除并返回最小元素(根据Less)也就是堆顶元素.
// 并进行下滤操作.
func (h *Heap[T]) PopElement() T {
	return Pop[T](h)
}

// Remove 从堆中移除并返回 index=i 的元素.
func (h *Heap[T]) Remove(i int) T {
	return Remove[T](h, i)
}

// Fix 在 index=i 的元素值改变后重新建立堆排序.
// 在更改 index=i 元素的值后调用 Fix(i).
// 相当于调用 Remove(i) 然后再插入新值.
func (h *Heap[T]) Fix(i int) {
	Fix[T](h, i)
}

// Replace 替换 index=i 的元素为 x.
// 并调用 Fix(i) 修复堆.
func (h *Heap[T]) Replace(i int, x T) {
	(*h)[i] = x
	h.Fix(i)
}
