package helpers

import farm "ants/pkg/farm"

type Queue struct {
	items []*farm.Room
}

func (q *Queue) Enqueue(item *farm.Room) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() *farm.Room {
	if len(q.items) == 0 {
		return nil // Queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
