package queue

import "sync"

type Queue struct {
	Values []int
	mutex  sync.Mutex
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(v int) {
	q.mutex.Lock()
	q.Values = append(q.Values, v)
	q.mutex.Unlock()
}

func (q *Queue) Dequeue() int {
	q.mutex.Lock()
	v := q.Values[0]
	q.Values = q.Values[1:]
	q.mutex.Unlock()
	return v

}

func (q *Queue) Size() int {
	return len(q.Values)
}
