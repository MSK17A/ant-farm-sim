package farm

func (farm *Farm) AntBFS() {
	queue := &Queue{}
	discovered := make(map[*Room]bool)
	visited := make(map[*Room]bool)

	queue.Enqueue(farm.end_room)
	room := queue.Dequeue()
	visited[room] = true // added to visited list

	for room != nil {
		for _, tunnel := range room.tunnels {
			if !visited[tunnel] && room.is_empty {
				//short_path_farm.AddRoom(tunnel.name, "normal", 0, 0)
				farm.AddShortTunnel(tunnel.name, room.name, false)
				visited[tunnel] = true
				queue.Enqueue(tunnel)
			}
		}
		discovered[room] = true
		room = queue.Dequeue()
	}
}
