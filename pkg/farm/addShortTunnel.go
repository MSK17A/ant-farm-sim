package farm

import "fmt"

/* Joins two rooms together */
func (farm *Farm) AddShortTunnel(from_room string, to_room string, bi_direction bool) {
	if farm.rooms[from_room] == nil || farm.rooms[to_room] == nil {
		fmt.Println("One room or two wasn't found!")
		return
	}
	if farm.rooms[from_room].Short_Tunnel_Exists(farm.rooms[to_room]) {
		return
	}
	if farm.rooms[to_room].Short_Tunnel_Exists(farm.rooms[from_room]) {
		return
	}
	if bi_direction {
		farm.rooms[from_room].short_tunnels.AddToList(farm.rooms[to_room])
		farm.rooms[to_room].short_tunnels.AddToList(farm.rooms[from_room])
	} else {
		farm.rooms[from_room].short_tunnels.AddToList(farm.rooms[to_room])
	}
}
