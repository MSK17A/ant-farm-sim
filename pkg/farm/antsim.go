package farm

import "fmt"

/* Starts the simulator with one step per call */
func (farm *Farm) AntSim_Step() {

	ants := farm.ants

	/* Loop throgh each ant */
	for ant_idx := range ants {
		ant := ants[ant_idx]
		alt_tun := farm.Find_Min_Path(ant)
		if (alt_tun.is_empty || alt_tun.end) && (!ant.discovered_rooms[alt_tun] && ant.moving) {
			ant.discovered_rooms[ant.room] = true
			ant.discovered_rooms[alt_tun] = true
			ant.room.is_empty = true // flag current room as empty
			ant.room = alt_tun       // go to next room
			alt_tun.is_empty = false // flag next room as not empty
			if ant.room.end {
				ant.moving = false
			}
		}
	}

}

/* Starts the simulator until all ants are at the end room */
func (farm *Farm) AntSim() {

	Step := 1
	for !farm.Ants_At_End() {
		fmt.Printf("\nAnts moves step %d:\n", Step)
		farm.AntSim_Step()
		farm.Print_Ants_Locations()
		Step++
	}

}

/* Custom iterations for Ant simulator */
func (farm *Farm) AntSim_Iter(iter int) {
	iteration := 0
	Step := 1
	for !farm.Ants_At_End() && iteration < iter {
		fmt.Printf("\nAnts moves step %d:\n", Step)
		farm.AntSim_Step()
		farm.Print_Ants_Locations()
		Step++
		iteration++
	}

}

/* This function will find the minimum path the ant can take */
func (farm *Farm) Find_Min_Path(ant *Ant) *Room {
	min := 9999
	temp := ant.room.tunnels.head.room

	tunnel := ant.room.tunnels.head
	for tunnel != nil {
		if farm.distances[tunnel.room] < min && (tunnel.room.is_empty || tunnel.room.end) && !ant.discovered_rooms[tunnel.room] && !tunnel.room.start {
			temp = tunnel.room
			min = farm.distances[temp]
		}
		tunnel = tunnel.next
	}
	return temp
}
