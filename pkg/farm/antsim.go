package farm

import (
	"fmt"
)

/* Starts the simulator with one step per call */
func (farm *Farm) AntSim_Step() {
	ants := farm.ants

	/* Loop throgh each ant */
	for ant_idx := range ants {

		ant := ants[ant_idx]
		alt_tun := farm.Find_Min_Path(ant)
		if check_moving_possiblity(ant, alt_tun) {
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
	prev_steps := 999999
	step_string := ""
	min_steps_string := ""

	for ant := 0; ant <= farm.number_of_ants; ant++ {

		step := 0
		step_string = ""
		for !farm.Ants_At_End() {
			step++
			//fmt.Printf("\nAnts moves step %d:\n", Step)
			farm.AntSim_Step()
			step_string += farm.Print_Ants_Locations()
		}
		if step < prev_steps {
			prev_steps = step
			min_steps_string = step_string
		}
		farm.Re_InitAnts()
	}

	fmt.Print(min_steps_string)
	fmt.Printf("\nSolution found with %d steps\n", prev_steps)
}

/* Custom iterations for Ant simulator */
func (farm *Farm) AntSim_Iter(iter int) {
	iteration := 0
	for ant := 0; ant <= farm.number_of_ants; ant++ {
		Step := 1
		for !farm.Ants_At_End() && iteration < iter {
			fmt.Printf("\nAnts moves step %d:\n", Step)
			farm.AntSim_Step()
			farm.Print_Ants_Locations()
			Step++
			iteration++
		}
	}
}

/* This function will find the minimum path the ant can take */
func (farm *Farm) Find_Min_Path(ant *Ant) *Room {
	min := 9999
	temp := ant.room.tunnels.head.room

	tunnel := ant.room.tunnels.head
	for tunnel != nil {
		if farm.distances[tunnel.room] < min && check_moving_possiblity(ant, tunnel.room) {
			temp = tunnel.room
			min = farm.distances[temp]
		}

		/* If these ant are last n ant then wait till the short path is available */
		tunnel = tunnel.next
	}
	return temp
}

func check_moving_possiblity(ant *Ant, tunnel *Room) bool {
	return (tunnel.is_empty || tunnel.end) && !ant.discovered_rooms[tunnel] && !tunnel.start && !ant.room.locked_tunnels[tunnel.name] && ant.moving
}

func (farm *Farm) Sort_first_move_tunnels() *LinkedRoomsList {
	tunnel := farm.start_room.tunnels.head
	already_sorted := make(map[string]bool)
	first_move_tunnels_sorted := &LinkedRoomsList{}
	min := 99999
	tunnels_outer_looper := farm.start_room.tunnels.head
	for tunnels_outer_looper != nil {
		for tunnel != nil && !already_sorted[tunnel.room.name] && farm.distances[tunnel.room] <= min {
			first_move_tunnels_sorted.AddToList(tunnel.room)
			min = farm.distances[tunnel.room]
			already_sorted[tunnel.room.name] = true
			tunnel = tunnel.next
		}
		min = 99999
		tunnels_outer_looper = tunnels_outer_looper.next
	}

	return first_move_tunnels_sorted
}

func (farm *Farm) Distribute_ant_starter() {
	//min_steps := 0
	sorted_first_tunnels := farm.Sort_first_move_tunnels()
	sorted_head := sorted_first_tunnels.head

	sorted_first_tunnels_steps := make(map[*Room]int)
	for sorted_head != nil {
		sorted_first_tunnels_steps[sorted_head.room] = farm.distances[sorted_head.room]
		sorted_head = sorted_head.next
	}
	sorted_head = sorted_first_tunnels.head

	for range sorted_first_tunnels_steps {
		if sorted_head.next == nil {
			break
		}
		/* loop through each ant */
		for ant_idx := range farm.ants {
			/* check if this tunnel distance steps is lower than the next */
			if sorted_first_tunnels_steps[sorted_head.room] <= sorted_first_tunnels_steps[sorted_head.next.room] && !farm.ants[ant_idx].self_start {
				/* Force the ant to move */
				farm.ants[ant_idx].self_start = true
				farm.ants[ant_idx].force_move_to_room = sorted_head.room
				sorted_first_tunnels_steps[sorted_head.room]++
			}
		}

		sorted_head = sorted_head.next
	}
}
