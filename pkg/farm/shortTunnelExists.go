package farm

import "fmt"

func (room *Room) Short_Tunnel_Exists(to_room *Room) bool {
	tunnel := room.short_tunnels.head

	for tunnel != nil {
		if tunnel.room.name == to_room.name {
			fmt.Printf("Short Tunnel %s -> %s already exist\n", room.name, to_room.name)
			return true
		}
		tunnel = tunnel.next
	}
	return false
}
