package farm

import "fmt"

func (farm *Farm) AntSim() {

	ant_rooms := farm.ants_rooms
	/* Loop throgh each ant */
	for ant_room_idx := range ant_rooms {
		tunnels := ant_rooms[ant_room_idx].tunnels
		/* Loop through each available tunnel */
		for tunnel_idx := range tunnels {
			/* If the tunneled room is empty, then move this ant to the next room */
			if tunnels[tunnel_idx].is_empty {
				ant_rooms[ant_room_idx].is_empty = true       // flag current room as empty
				ant_rooms[ant_room_idx] = tunnels[tunnel_idx] // go to next room
				tunnels[tunnel_idx].is_empty = false          // flag next room as not empty
				break
			}
		}
	}

	for i, ant_room := range farm.ants_rooms {
		fmt.Printf("ant %d in %s\n", i, ant_room.name)
	}
}
