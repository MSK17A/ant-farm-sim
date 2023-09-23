package farm

import "fmt"

func (room *Room) Tunnel_Exists(to_room *Room) bool {
	for _, tunnel := range room.tunnels {
		if tunnel.name == to_room.name {
			fmt.Printf("Tunnel %s -> %s already exist\n", room.name, to_room.name)
			return true
		}
	}
	return false
}
