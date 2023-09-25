package farm

func (farm *Farm) FreeShortTunnels() {
	for idx := range farm.rooms {
		room := farm.rooms[idx]
		for tun_idx := range room.short_tunnels {
			room.short_tunnels[tun_idx] = &Room{
				short_tunnels: nil,
			}
		}
	}
}
