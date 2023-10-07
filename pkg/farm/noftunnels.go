package farm

func (farm *Farm) Number_of_Tunnels(ant *Ant) int {

	counter := 0

	tunnel := ant.room.tunnels.head

	for tunnel != nil {
		if (farm.distances[tunnel.room] <= farm.distances[ant.room]) && !ant.discovered_rooms[tunnel.room] {
			counter++
		}
		tunnel = tunnel.next
	}

	return counter
}
