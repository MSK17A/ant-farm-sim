package farm

func (farm *Farm) AntBFS() *Farm {
	queue := &Queue{}
	var short_path_farm Farm
	short_path_farm.InitFarm()
	discovered := make(map[*Room]bool)
	visited := make(map[*Room]bool)

	queue.Enqueue(farm.end_room)
	room := queue.Dequeue()
	short_path_farm.AddRoom(room.name, "end", 0, 0) // Construct a new room
	visited[room] = true                            // added to visited list

	for room != nil {
		for _, tunnel := range room.tunnels {
			if !visited[tunnel] && tunnel.is_empty {
				short_path_farm.AddRoom(tunnel.name, "normal", 0, 0)
				short_path_farm.AddTunnel(tunnel.name, room.name, false)
				visited[tunnel] = true
				queue.Enqueue(tunnel)
			}
		}
		discovered[room] = true
		room = queue.Dequeue()
	}

	return &short_path_farm
}
