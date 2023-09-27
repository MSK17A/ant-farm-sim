package farm

import "fmt"

/* Starts the simulator with one step per call */
func (farm *Farm) AntSim_Step() {
	ants := farm.ants

	/* Loop throgh each ant */
	for ant_idx := range ants {

		ant := ants[ant_idx]
		alt_tun := farm.Find_Min_Path(ant, ant_idx)
		if (alt_tun.is_empty || alt_tun.end) && !ant.discovered_rooms[alt_tun] && ant.moving && !ant.room.locked_tunnels[alt_tun.name] {
			//ant.discovered_rooms[alt_tun] = true  // remember the next room
			ant.room.locked_tunnels[alt_tun.name] = true // Lock the tunnel from beign used by other ant until step is finished
			ant.discovered_rooms[ant.room] = true        // remember the current room
			ant.room.is_empty = true                     // flag current room as empty
			ant.room = alt_tun                           // go to next room
			alt_tun.is_empty = false                     // flag next room as not empty
			if ant.room.end {
				ant.moving = false

			}
			ant.moved = true

		}
	}
	farm.Unlock_Locked_Tunnel()
}

/* Starts the simulator until all ants are at the end room */
func (farm *Farm) AntSim() {

	Step := 1
	for !farm.Ants_At_End() {
		//fmt.Printf("\nAnts moves step %d:\n", Step)
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
func (farm *Farm) Find_Min_Path(ant *Ant, ant_idx int) *Room {
	min := 9999
	temp := ant.room.tunnels.head.room

	tunnel := ant.room.tunnels.head
	for tunnel != nil {
		if farm.distances[tunnel.room] < min && (tunnel.room.is_empty || tunnel.room.end) && !ant.discovered_rooms[tunnel.room] && !tunnel.room.start && !ant.room.locked_tunnels[tunnel.room.name] {
			temp = tunnel.room
			min = farm.distances[temp]
		}
		/* If this ant is last ant then wait till the short path is available */
		if ant_idx > farm.number_of_ants-(Number_of_Tunnels(ant.room.tunnels)) && farm.distances[tunnel.room] <= min {
			temp = tunnel.room
			min = farm.distances[temp]
		}
		tunnel = tunnel.next
	}
	return temp
}

func (farm *Farm) Unlock_Locked_Tunnel() {
	ants := farm.ants

	for ants_idx := range ants {
		ant := ants[ants_idx]

		for locked_tunnel_idx := range ant.room.locked_tunnels {
			ant.room.locked_tunnels[locked_tunnel_idx] = false
		}
	}
}
