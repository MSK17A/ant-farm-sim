package farm

func (farm *Farm) AntBFS() {
	queue := &Queue{}

	//discovered := make(map[*Room]bool)
	visited := make(map[*Room]bool)
	//not_visited := make(map[*Room]bool)

	queue.Enqueue(farm.end_room)
	room := queue.Dequeue()

	for room != nil {
		visited[room] = true
	}

}
