package farm

import (
	"fmt"
	"sort"
)

var (
	one_ant_moved = false
)

/* Starts the simulator with one step per call */
func (farm *Farm) AntSim_Step() {
	ants_to_work_on := make([]*Ant, farm.number_of_ants)
	copy(ants_to_work_on, farm.ants)

	/* Ants with less tunnels will go first */
	sort.SliceStable(ants_to_work_on, func(i, j int) bool {
		return farm.same_distance_tunnels(ants_to_work_on[i]) < farm.same_distance_tunnels(ants_to_work_on[j])
	})

	/* Loop throgh each ant */
	for ant_idx := 0; ant_idx < len(ants_to_work_on); ant_idx++ {
		alt_tun := farm.Find_Min_Path(ants_to_work_on[ant_idx])

		/* Will cause the ant to hold in one place so that in the next step when the tunnel is empty, it will take it */
		/* Distance of steps is increased because one ant has taken this path, or will take it in the next step */
		if alt_tun != farm.end_room || ants_to_work_on[ant_idx].room.start {
			farm.distances[alt_tun]++
		}
		if check_moving_possiblity(ants_to_work_on[ant_idx], alt_tun) {

			ants_to_work_on[ant_idx].room.locked_tunnels[alt_tun.name] = true               // Lock the tunnel from beign used by other ant until step is finished
			ants_to_work_on[ant_idx].discovered_rooms[ants_to_work_on[ant_idx].room] = true // remember the current room
			ants_to_work_on[ant_idx].room.is_empty = true                                   // flag current room as empty
			ants_to_work_on[ant_idx].room = alt_tun                                         // go to next room
			alt_tun.is_empty = false                                                        // flag next room as not empty
			if ants_to_work_on[ant_idx].room.end {
				ants_to_work_on[ant_idx].moving = false
			}
			ants_to_work_on[ant_idx].moved = true
			one_ant_moved = true
		}

		if ant_idx == len(ants_to_work_on)-1 && one_ant_moved {
			ant_idx = 0
			one_ant_moved = false
		}

	}

	farm.PrintDistances()
	farm.Unlock_Locked_Tunnel()
	farm.AntBFS()

}

/* Starts the simulator until all ants are at the end room */
func (farm *Farm) AntSim() {

	step_string := ""
	step := 0
	for !farm.Ants_At_End() {
		step++
		//fmt.Printf("\nAnts moves step %d:\n", Step)
		farm.AntSim_Step()
		step_string += farm.Print_Ants_Locations()
	}

	fmt.Print(step_string)
	fmt.Printf("\nSolution found with %d steps\n", step)

	/*fmt.Print(step_string)
	fmt.Printf("\nSolution found with %d steps\n", step)*/
}

/* Custom iterations for Ant simulator */
func (farm *Farm) AntSim_Iter(iter int) {
	iteration := 0
	step_string := ""
	step := 0
	for ant := 0; ant <= farm.number_of_ants; ant++ {
		for !farm.Ants_At_End() && iteration < iter {
			step++
			fmt.Printf("\nAnts moves step %d:\n", step)
			farm.AntSim_Step()
			step_string += farm.Print_Ants_Locations()
			iteration++
		}
	}
	fmt.Print(step_string)
	fmt.Printf("\nSolution found with %d steps\n", step)
}

/* This function will find the minimum path the ant can take */
func (farm *Farm) Find_Min_Path(ant *Ant) *Room {
	min := 9999
	temp := ant.room.tunnels.head.room

	tunnel := ant.room.tunnels.head
	for tunnel != nil {

		if farm.distances[tunnel.room] <= min && !ant.discovered_rooms[tunnel.room] {
			temp = tunnel.room
			min = farm.distances[temp]

		}

		tunnel = tunnel.next
	}
	return temp
}

/* Checks the possibility of the ant to take a specified tunnel */
func check_moving_possiblity(ant *Ant, tunnel *Room) bool {
	return (tunnel.is_empty || tunnel.end) && !ant.discovered_rooms[tunnel] && !tunnel.start && !ant.room.locked_tunnels[tunnel.name] && ant.moving && !tunnel.dead_end && !ant.moved
}

/* Counts the tunnels that have same distance */
func (farm *Farm) same_distance_tunnels(ant *Ant) int {
	tunnel := ant.room.tunnels.head
	similar_distance_count := 0

	for tunnel != nil {
		compare_tunnel := tunnel.next
		for compare_tunnel != nil {
			if farm.distances[tunnel.room] == farm.distances[compare_tunnel.room] && (farm.distances[tunnel.room] <= farm.distances[ant.room]) && !ant.discovered_rooms[tunnel.room] {
				similar_distance_count++
			}
			compare_tunnel = compare_tunnel.next
		}
		tunnel = tunnel.next
	}

	return similar_distance_count
}
