package utils

import "container/list"

type Queue struct {
	lst *list.List 
}

func New() *Queue {
	return &Queue{ lst: list.New() }
}

func (q *Queue) Add(el any) *list.Element {
	return q.lst.PushBack(el)
}

func (q *Queue) Pop() *list.Element {
	return q.lst.Front()
}
