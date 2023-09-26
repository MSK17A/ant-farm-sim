package farm

import "fmt"

func (farm *Farm) AntSim_Step() {

	ants := farm.ants

	/* Loop throgh each ant */
	for ant_idx := range ants {

		/* Loop through each available tunnel */
		for _, tunnels := range ants[ant_idx].room.tunnels {
			/* If the tunneled room is empty, then move this ant to the next room */
			if (tunnels.is_empty || tunnels.end) && (!ants[ant_idx].discovered_rooms[tunnels] && ants[ant_idx].moving) {
				ants[ant_idx].discovered_rooms[tunnels] = true
				ants[ant_idx].room.is_empty = true // flag current room as empty
				ants[ant_idx].room = tunnels       // go to next room
				tunnels.is_empty = false           // flag next room as not empty
				if ants[ant_idx].room.end {
					ants[ant_idx].moving = false
				}
				break
			}
			tunnel_alt := farm.Find_Min_Path(ants[ant_idx].room.tunnels)
			ants[ant_idx].discovered_rooms[tunnels] = true
			ants[ant_idx].room.is_empty = true // flag current room as empty
			ants[ant_idx].room = tunnel_alt    // go to next room
			tunnels.is_empty = false           // flag next room as not empty
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

func (farm *Farm) AntSim_Iter(iter int) {
	iteration := 0
	Step := 1
	for !farm.Ants_Stuck() && iteration < iter {
		fmt.Printf("\nAnts moves step %d:\n", Step)
		farm.AntSim_Step()
		farm.Print_Ants_Locations()
		Step++
		iteration++
	}

}

func (farm *Farm) Find_Min_Path(tunnels []*Room) *Room {
	min := 9999

	for tunnel_idx := range tunnels {
		if farm.distances[tunnels[tunnel_idx]] < min && tunnels[tunnel_idx].is_empty {
			return tunnels[tunnel_idx]
		}
	}
	return nil
}
