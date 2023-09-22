package farm

import "fmt"

/* Joins two rooms together */
func (farm *Farm) AddTunnel(from_room string, to_room string) {
	if farm.Rooms[from_room] == nil || farm.Rooms[to_room] == nil {
		fmt.Println("One room or two wasn't found!")
		return
	}
	if farm.Rooms[from_room].Tunnel_Exists(farm.Rooms[to_room]) {
		return
	}
	if farm.Rooms[to_room].Tunnel_Exists(farm.Rooms[from_room]) {
		return
	}
	farm.Rooms[from_room].tunnels = append(farm.Rooms[from_room].tunnels, farm.Rooms[to_room])
}
