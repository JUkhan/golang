package ds

import (
	"sync"
)

//this is queue item type

type Item interface {
	//ToString() string
}

type Queue struct {
	items  []Item
	locker sync.RWMutex
}

func (q *Queue) New() *Queue {
	q.items = []Item{}
	return q
}
func (q *Queue) Enqueue(data Item) {
	q.locker.Lock()
	q.items = append(q.items, data)
	q.locker.Unlock()
}
func (q *Queue) Dequeue() (res Item) {
	q.locker.Lock()
	res = q.items[0]
	q.items = q.items[1:len(q.items)]
	q.locker.Unlock()
	return
}
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
func (q *Queue) Front() Item {
	return q.items[0]
}
