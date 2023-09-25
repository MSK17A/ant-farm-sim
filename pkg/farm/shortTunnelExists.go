package farm

import "fmt"

func (room *Room) Short_Tunnel_Exists(to_room *Room) bool {
	for _, tunnel := range room.short_tunnels {
		if tunnel.name == to_room.name {
			fmt.Printf("Short Tunnel %s -> %s already exist\n", room.name, to_room.name)
			return true
		}
	}
	return false
}
