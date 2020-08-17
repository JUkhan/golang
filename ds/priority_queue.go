package ds

import "sync"

type PriorityQueue struct {
	items  []Item
	locker sync.RWMutex

	comparato func(parentItem, childItem Item) bool
}

func NewPriorityQueue(comparato func(parentItem, childItem Item) bool) *PriorityQueue {
	q := PriorityQueue{comparato: comparato}
	q.items = []Item{nil}
	return &q
}
func (q *PriorityQueue) Heapify(arr []Item) {
	q.locker.Lock()
	q.items = append(arr, nil)
	l := len(q.items)
	q.items[0], q.items[l-1] = q.items[l-1], q.items[0]
	for j := l - 1; j > 0; j-- {
		if j*2 > l {
			continue
		}
		q.adjustDownword(j, l)
	}
	q.locker.Unlock()

}
func (q *PriorityQueue) Enqueue(data Item) {
	q.locker.Lock()
	q.items = append(q.items, data)
	l := len(q.items)
	if l > 2 {
		for j := l - 1; j > 1; {
			p := j / 2
			if q.comparato(q.items[p], q.items[j]) {
				q.items[p], q.items[j] = q.items[j], q.items[p]
			}
			j = p
		}
	}

	q.locker.Unlock()
}
func (q *PriorityQueue) adjustDownword(i, l int) {
	for i < l {
		lc := 2 * i
		rc := 2*i + 1
		if lc < l && rc < l {
			if q.comparato(q.items[lc], q.items[rc]) {
				lc = rc
			}
			if q.comparato(q.items[i], q.items[lc]) {
				q.items[i], q.items[lc] = q.items[lc], q.items[i]
			}

		} else if lc < l {
			if q.comparato(q.items[i], q.items[lc]) {
				q.items[i], q.items[lc] = q.items[lc], q.items[i]
			}

		}
		i = lc
	}
}
func (q *PriorityQueue) Dequeue() (res Item) {
	q.locker.Lock()
	l := len((q.items))
	if l == 1 {
		return
	}
	res = q.items[1]
	if l > 2 {
		q.items[1] = q.items[l-1]
		q.items = q.items[:l-1]
		l--
		q.adjustDownword(1, l)

	} else {
		q.items = []Item{nil}
	}
	q.locker.Unlock()
	return
}
func (q *PriorityQueue) IsEmpty() bool {
	return len(q.items) == 1
}
func (q *PriorityQueue) Front() Item {
	return q.items[1]
}
func (q *PriorityQueue) GetTree() []Item {
	return q.items
}
