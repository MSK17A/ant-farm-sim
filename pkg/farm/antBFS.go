package farm

import helpers "ants/pkg/helpers"

func (farm *Farm) AntBFS() {
	var queue helpers.Queue

	//var discovered map[*Room]bool
	var visited map[*Room]bool
	//var not_visited map[*Room]bool

	visited = make(map[*Room]bool)

	queue.Enqueue(farm.end_room)
	room := queue.Dequeue()

	for room != nil {
		visited[room] = true
	}

}
