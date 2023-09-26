package farm

func (farm *Farm) AntBFS() {
	queue := &Queue{}
	discovered := make(map[*Room]bool)
	visited := make(map[*Room]bool)

	queue.Enqueue(farm.end_room)
	room := queue.Dequeue()
	farm.distances[room] = 1
	visited[room] = true // added to visited list

	for room != nil {
		tunnel := room.tunnels.head
		for tunnel != nil {
			if !discovered[tunnel.room] {
				if farm.distances[room]+1 <= farm.distances[tunnel.room] {
					farm.distances[tunnel.room] = farm.distances[room] + 1
					queue.Enqueue(tunnel.room)
				}
				visited[tunnel.room] = true
				queue.Enqueue(tunnel.room)
			}
			tunnel = tunnel.next
		}
		discovered[room] = true
		room = queue.Dequeue()
	}
}
