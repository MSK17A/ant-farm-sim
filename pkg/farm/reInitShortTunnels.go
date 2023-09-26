package farm

func (farm *Farm) ReInitShortTunnels() {

	for room_idx := range farm.rooms {
		farm.rooms[room_idx].short_tunnels = &LinkedRoomsList{}
	}
}
