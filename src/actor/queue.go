package actor

import (
	"container/list"
	"sync"
)

type Queue struct {
	l    list.List
	lock sync.Mutex
}

func (q *Queue) push(v interface{}) {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.l.PushBack(v)
}

func (q *Queue) pop() interface{} {
	defer q.lock.Unlock()
	q.lock.Lock()
	e := q.l.Front()
	if e == nil {
		return nil
	}
	q.l.Remove(e)
	return e.Value
}
