package farm

func (farm *Farm) AntBFS() {
	queue := &Queue{}

	//var discovered map[*Room]bool
	//var visited map[interface{}]bool
	//var not_visited map[*Room]bool

	visited := make(map[*Room]bool)

	queue.Enqueue(farm.end_room)
	room := queue.Dequeue()

	for room != nil {
		visited[room] = true
	}

}
