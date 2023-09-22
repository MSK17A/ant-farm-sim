package farm

func (farm *Farm) AntSim() {
	/* Loop throgh each ant */
	for _, ant_room := range farm.rooms {
		/* Loop through each available tunnel */
		for _, tunnel := range ant_room.tunnels {
			/* If the tunneled room is empty, then move this ant to the next room */
			if tunnel.is_empty {
				ant_room.is_empty = true // flag current room as empty
				ant_room = tunnel        // go to next room
				tunnel.is_empty = false  // flag next room as not empty
				break
			}
		}
	}
}
