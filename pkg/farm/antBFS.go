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
		for tunnel_idx := range room.tunnels {
			tunnel := room.tunnels[tunnel_idx]

			if !discovered[tunnel] {
				if farm.distances[room]+1 <= farm.distances[tunnel] {
					farm.distances[tunnel] = farm.distances[room] + 1
					queue.Enqueue(tunnel)
				}
				visited[tunnel] = true
				queue.Enqueue(tunnel)
			}

		}
		discovered[room] = true
		room = queue.Dequeue()
	}
}
