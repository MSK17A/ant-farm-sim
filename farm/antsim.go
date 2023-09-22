package farm

import "fmt"

func (farm *Farm) AntSim_Step() {

	ants := farm.ants
	/* Loop throgh each ant */
	for ant_idx := range ants {
		tunnels := ants[ant_idx].room.tunnels

		/* Loop through each available tunnel */
		for tunnel_idx := range tunnels {
			tunnel := tunnels[tunnel_idx]
			/* If the tunneled room is empty, then move this ant to the next room */
			if (tunnel.is_empty || tunnel.end) && !ants[ant_idx].discovered_rooms[tunnel.name] && ants[ant_idx].moving {
				ants[ant_idx].discovered_rooms[tunnel.name] = true
				ants[ant_idx].room.is_empty = true // flag current room as empty
				ants[ant_idx].room = tunnel        // go to next room
				tunnel.is_empty = false            // flag next room as not empty
				if tunnel.end {
					ants[ant_idx].moving = false
				}
				break
			}
		}
	}
}

func (farm *Farm) AntSim() {
	Step := 1
	for !farm.Ants_Stuck() {
		fmt.Printf("\nAnts moves step %d:\n", Step)
		farm.AntSim_Step()
		farm.Print_Ants_Locations()
		Step++
	}

}
