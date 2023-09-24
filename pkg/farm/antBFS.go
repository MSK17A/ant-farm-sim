package farm

import helpers "ants/helpers"

func (farm *Farm) AntBFS() {
	queue := &helpers.Queue{}

	//var discovered map[*Room]bool
	//var visited map[interface{}]bool
	//var not_visited map[*Room]bool

	visited := make(map[interface{}]bool)

	queue.Enqueue(farm.end_room)
	room := queue.Dequeue()

	for room != nil {
		visited[room] = true
	}

}
