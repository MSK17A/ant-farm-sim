package helpers

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil // Queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
