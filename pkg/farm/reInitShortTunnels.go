package farm

func (farm *Farm) ReInitShortTunnels() {

	for room_idx := range farm.rooms {
		if farm.rooms[room_idx].is_empty {
			farm.rooms[room_idx].short_tunnels = &LinkedRoomsList{}
		}
	}
}
