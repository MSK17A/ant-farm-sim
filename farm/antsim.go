package farm

import "fmt"

func (farm *Farm) AntSim_Step() {

	ant_rooms := farm.ants_rooms
	/* Loop throgh each ant */
	for ant_room_idx := range ant_rooms {
		tunnels := ant_rooms[ant_room_idx].tunnels

		/* Loop through each available tunnel */
		for tunnel_idx := range tunnels {
			tunnel := tunnels[tunnel_idx]

			/* If the tunneled room is empty, then move this ant to the next room */
			if tunnel.is_empty {
				ant_rooms[ant_room_idx].is_empty = true // flag current room as empty
				ant_rooms[ant_room_idx] = tunnel        // go to next room
				tunnel.is_empty = false                 // flag next room as not empty
				break
			}
		}
	}

	for i, ant_room := range farm.ants_rooms {
		fmt.Printf("ant %d in %s\n", i+1, ant_room.name)
	}
}