package farm

import "fmt"

func (farm *Farm) AntSim_Step() {

	ants := farm.ants

	/* Loop throgh each ant */
	for ant_idx := range ants {

		alt_tun := farm.Find_Min_Path(ants[ant_idx])
		if (alt_tun.is_empty || alt_tun.end) && (!ants[ant_idx].discovered_rooms[alt_tun] && ants[ant_idx].moving) {
			ants[ant_idx].discovered_rooms[alt_tun] = true
			ants[ant_idx].room.is_empty = true // flag current room as empty
			ants[ant_idx].room = alt_tun       // go to next room
			alt_tun.is_empty = false           // flag next room as not empty
			if ants[ant_idx].room.end {
				ants[ant_idx].moving = false
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

func (farm *Farm) Find_Min_Path(ant *Ant) *Room {
	min := 9999
	temp := ant.room.tunnels[0]

	tunnels := ant.room.tunnels
	for tunnel_idx := range tunnels {
		if farm.distances[tunnels[tunnel_idx]] < min && (tunnels[tunnel_idx].is_empty || tunnels[tunnel_idx].end) && !ant.discovered_rooms[tunnels[tunnel_idx]] {
			temp = tunnels[tunnel_idx]
			min = farm.distances[temp]
		}
	}
	return temp
}
