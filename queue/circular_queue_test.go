package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue_Enqueue(t *testing.T) {
	type testCase[T any] struct {
		name    string
		cq      *CircularQueue[T]
		val     T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name:    "normal",
			cq:      NewCircularQueue[int](1),
			val:     1,
			wantErr: nil,
		},
		{
			name: "queue_is_full",
			cq: &CircularQueue[int]{
				capacity: 2,
				head:     0,
				tail:     1,
				data:     []int{1, 0},
			},
			val:     1,
			wantErr: ErrFullQueue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantErr, tt.cq.Enqueue(tt.val))
		})
	}
}

func TestCircularQueue_Dequeue(t *testing.T) {
	type testCase[T any] struct {
		name    string
		cq      *CircularQueue[T]
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "normal",
			cq: &CircularQueue[int]{
				capacity: 2,
				head:     0,
				tail:     1,
				data:     []int{9, 0},
			},
			want:    9,
			wantErr: nil,
		},
		{
			name: "queue_is_empty",
			cq: &CircularQueue[int]{
				capacity: 2,
				head:     1,
				tail:     1,
				data:     []int{9, 0},
			},
			want:    0,
			wantErr: ErrEmptyQueue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cq.Dequeue()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCircularQueue_Peek(t *testing.T) {
	type testCase[T any] struct {
		name    string
		c       *CircularQueue[T]
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "normal",
			c: &CircularQueue[int]{
				capacity: 2,
				head:     0,
				tail:     1,
				data:     []int{5, 0},
			},
			want:    5,
			wantErr: nil,
		},
		{
			name: "queue_is_empty",
			c: &CircularQueue[int]{
				capacity: 2,
				head:     1,
				tail:     1,
				data:     []int{5, 0},
			},
			want:    0,
			wantErr: ErrEmptyQueue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Peek()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCircularQueue_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name string
		c    *CircularQueue[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "is_empty",
			c:    NewCircularQueue[int](1),
			want: true,
		},
		{
			name: "not_empty",
			c: &CircularQueue[int]{
				capacity: 2,
				head:     0,
				tail:     1,
				data:     []int{1, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.c.IsEmpty())
		})
	}
}

func TestCircularQueue_IsFull(t *testing.T) {
	type testCase[T any] struct {
		name string
		c    *CircularQueue[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "is_full",
			c: &CircularQueue[int]{
				capacity: 2,
				head:     0,
				tail:     1,
				data:     []int{1, 0},
			},
			want: true,
		},
		{
			name: "not_full",
			c: &CircularQueue[int]{
				capacity: 2,
				head:     0,
				tail:     0,
				data:     []int{0, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.c.IsFull())
		})
	}
}

func TestCircularQueueLifecycle(t *testing.T) {
	cq := NewCircularQueue[int](2)
	type testCase[T any] struct {
		name     string
		op       func() (int, bool, error) // operation 操作
		want     T
		wantBool bool
		wantErr  error
	}
	tests := []testCase[int]{
		{
			name: "is_empty",
			op: func() (int, bool, error) {
				return 0, cq.IsEmpty(), nil
			},
			wantBool: true,
		},
		{
			name: "enqueue",
			op: func() (int, bool, error) {
				err := cq.Enqueue(6)
				return 0, false, err
			},
		},
		{
			name: "not_empty",
			op: func() (int, bool, error) {
				return 0, cq.IsEmpty(), nil
			},
		},
		{
			name: "not_full",
			op: func() (int, bool, error) {
				return 0, cq.IsFull(), nil
			},
		},
		{
			name: "enqueue_another",
			op: func() (int, bool, error) {
				err := cq.Enqueue(7)
				return 0, false, err
			},
		},
		{
			name: "is_full",
			op: func() (int, bool, error) {
				return 0, cq.IsFull(), nil
			},
			wantBool: true,
		},
		{
			name: "enqueue_failed_queue_is_full",
			op: func() (int, bool, error) {
				err := cq.Enqueue(8)
				return 0, false, err
			},
			wantErr: ErrFullQueue,
		},
		{
			name: "peek",
			op: func() (int, bool, error) {
				peek, err := cq.Peek()
				return peek, false, err
			},
			want: 6,
		},
		{
			name: "dequeue",
			op: func() (int, bool, error) {
				dequeue, err := cq.Dequeue()
				return dequeue, false, err
			},
			want: 6,
		},
		{
			name: "peek_another",
			op: func() (int, bool, error) {
				peek, err := cq.Peek()
				return peek, false, err
			},
			want: 7,
		},
		{
			name: "enqueue_extra_one",
			op: func() (int, bool, error) {
				err := cq.Enqueue(8)
				return 0, false, err
			},
		},
		{
			name: "dequeue_another",
			op: func() (int, bool, error) {
				dequeue, err := cq.Dequeue()
				return dequeue, false, err
			},
			want: 7,
		},
		{
			name: "peek_extra_one",
			op: func() (int, bool, error) {
				peek, err := cq.Peek()
				return peek, false, err
			},
			want: 8,
		},
		{
			name: "dequeue_extra_one",
			op: func() (int, bool, error) {
				dequeue, err := cq.Dequeue()
				return dequeue, false, err
			},
			want: 8,
		},
		{
			name: "queue_is_empty",
			op: func() (int, bool, error) {
				return 0, cq.IsEmpty(), nil
			},
			wantBool: true,
		},
		{
			name: "dequeue_failed_queue_is_empty",
			op: func() (int, bool, error) {
				dequeue, err := cq.Dequeue()
				return dequeue, false, err
			},
			want:    0,
			wantErr: ErrEmptyQueue,
		},
	}
	for _, tt := range tests {
		got1, got2, err := tt.op()
		assert.Equalf(t, tt.wantErr, err, "%s: want %v; but got %v", tt.name, tt.wantErr, err)
		assert.Equalf(t, tt.want, got1, "%s: want %v; but got %v", tt.name, tt.want, got1)
		assert.Equalf(t, tt.wantBool, got2, "%s: want %v; but got %v", tt.name, tt.wantBool, got2)
	}
}

func TestNewCircularQueue(t *testing.T) {
	type testCase[T any] struct {
		name      string
		capacity  int
		want      *CircularQueue[T]
		wantPanic bool
	}
	tests := []testCase[int]{
		{
			name:     "normal",
			capacity: 1,
			want: &CircularQueue[int]{
				capacity: 2,
				head:     0,
				tail:     0,
				data:     make([]int, 2),
			},
			wantPanic: false,
		},
		{
			name:      "capacity_less_than_1",
			capacity:  0,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := NewCircularQueue[int]
			if tt.wantPanic {
				assert.Panics(t, func() { fn(tt.capacity) })
				return
			}
			assert.Equal(t, tt.want, fn(tt.capacity))
		})
	}
}
