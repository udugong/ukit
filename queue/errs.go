package queue

import "errors"

var (
	ErrEmptyQueue = errors.New("ukit: 队列为空")
	ErrFullQueue  = errors.New("ukit: 队列已满")
)
