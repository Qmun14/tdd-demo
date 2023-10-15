package queue

import (
	"fmt"
	"sort"
)

var (
	ErrQueueIsEmpty = fmt.Errorf("queue is Empty")
)

func New() *Queue {
	return &Queue{}
}

func NewPriorityQueue() *Queue {
	return &Queue{isPriorityQueue: true}
}

type Queue struct {
	data            []int
	isPriorityQueue bool
}

func (q Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Push(num int) {
	q.data = append(q.data, num)
	if q.isPriorityQueue {
		sort.Ints(q.data)
	}
}

func (q Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return -1, ErrQueueIsEmpty
	}
	return q.data[0], nil
}

func (q *Queue) Pop() (int, error) {
	val, err := q.Peek()
	if err != nil {
		return val, ErrQueueIsEmpty
	}
	q.data = q.data[1:]
	return val, nil
}
