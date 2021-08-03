package queue

import "sync"

type Queue struct {
	mt  sync.Mutex
	val []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		val: make([]interface{}, 0),
	}
}

func (q *Queue) Enqueue(el interface{}) {
	q.mt.Lock()
	defer q.mt.Unlock()

	q.val = append(q.val, el)
}

func (q *Queue) Dequeue() interface{} {
	q.mt.Lock()
	defer q.mt.Unlock()

	if q.Count() == 0 {
		return nil
	}

	val := q.val[0]
	q.val = q.val[1:]
	return val
}

func (q *Queue) Peek() interface{} {
	q.mt.Lock()
	defer q.mt.Unlock()

	if q.Count() == 0 {
		return nil
	}

	val := q.val[0]
	return val
}

func (q *Queue) Count() int {
	return len(q.val)
}
