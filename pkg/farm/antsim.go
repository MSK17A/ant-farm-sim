package farm

import (
	"fmt"
)

/* used for ants */
type antQueue struct {
	items []*Ant
}

func (q *antQueue) enqueue(item *Ant) {
	q.items = append(q.items, item)
}

func (q *antQueue) dequeue() *Ant {
	if len(q.items) == 0 {
		return nil // Queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

var (
	two_tunnels_hold = make(map[*Ant]bool)
)

/* Starts the simulator with one step per call */
func (farm *Farm) AntSim_Step(toggler bool) {
	ants := farm.ants
	ant_queue := &antQueue{}
	for ant_idx := range ants {
		ant_queue.enqueue(ants[ant_idx])
		two_tunnels_hold[farm.ants[ant_idx]] = false
	}

	/* Loop throgh each ant */
	ant := ant_queue.dequeue()
	for ant != nil {
		alt_tun := farm.Find_Min_Path(ant, toggler)
		/*if ant.self_start {
			alt_tun = ant.force_move_to_room
			ant.self_start = false
		}*/
		if check_moving_possiblity(ant, alt_tun) {
			if farm.same_distance_tunnels(ant) && !ant.room.start && !ant.room.end && !two_tunnels_hold[ant] {
				fmt.Printf("Room: %s\n", ant.room.name)
				ant_queue.enqueue(ant)
				two_tunnels_hold[ant] = true
				//ant = ant_queue.dequeue()
				continue
			}
			farm.distances[alt_tun]++
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

		} else {
			/* Will cause the ant to hold in one place so that in the next step when the tunnel is empty, it will take it */
			farm.distances[alt_tun]++
			ant.room.locked_tunnels[alt_tun.name] = true
		}
		ant = ant_queue.dequeue()
	}
	//farm.PrintDistances()
	farm.Unlock_Locked_Tunnel()
	farm.AntBFS()

}

/* Starts the simulator until all ants are at the end room */
func (farm *Farm) AntSim() {
	toggler_on_steps := 0
	toggler_off_steps := 0
	toggler_on_string := ""
	toggler_off_string := ""

	/*step_string := ""
	step := 0*/
	for !farm.Ants_At_End() {
		toggler_on_steps++
		//fmt.Printf("\nAnts moves step %d:\n", Step)
		farm.AntSim_Step(true)
		toggler_on_string += farm.Print_Ants_Locations()
	}

	farm.Re_InitAnts()
	for !farm.Ants_At_End() {
		toggler_off_steps++
		//fmt.Printf("\nAnts moves step %d:\n", Step)
		farm.AntSim_Step(false)
		toggler_off_string += farm.Print_Ants_Locations()
	}

	if toggler_off_steps == toggler_on_steps {
		fmt.Print(toggler_on_string)
		fmt.Printf("\nSolution found with %d steps\n", toggler_on_steps)
	} else if toggler_off_steps < toggler_on_steps {
		fmt.Print(toggler_off_string)
		fmt.Printf("\nSolution found with %d steps\n", toggler_off_steps)
	} else {
		fmt.Print(toggler_on_string)
		fmt.Printf("\nSolution found with %d steps\n", toggler_on_steps)
	}

	/*fmt.Print(step_string)
	fmt.Printf("\nSolution found with %d steps\n", step)*/
}

/* Custom iterations for Ant simulator */
func (farm *Farm) AntSim_Iter(iter int) {
	iteration := 0
	for ant := 0; ant <= farm.number_of_ants; ant++ {
		Step := 1
		for !farm.Ants_At_End() && iteration < iter {
			fmt.Printf("\nAnts moves step %d:\n", Step)
			farm.AntSim_Step(true)
			farm.Print_Ants_Locations()
			Step++
			iteration++
		}
	}
}

/* This function will find the minimum path the ant can take */
func (farm *Farm) Find_Min_Path(ant *Ant, toggler bool) *Room {
	min := 9999
	temp := ant.room.tunnels.head.room

	tunnel := ant.room.tunnels.head
	for tunnel != nil {
		if toggler {
			if farm.distances[tunnel.room] <= min && !ant.discovered_rooms[tunnel.room] {
				temp = tunnel.room
				min = farm.distances[temp]
			}
		} else {
			if farm.distances[tunnel.room] < min && !ant.discovered_rooms[tunnel.room] {
				temp = tunnel.room
				min = farm.distances[temp]
			}
		}

		/* If these ant are last n ant then wait till the short path is available */
		tunnel = tunnel.next
	}
	return temp
}

func check_moving_possiblity(ant *Ant, tunnel *Room) bool {
	return (tunnel.is_empty || tunnel.end) && !ant.discovered_rooms[tunnel] && !tunnel.start && !ant.room.locked_tunnels[tunnel.name] && ant.moving && !tunnel.dead_end
}

/*func (farm *Farm) Sort_first_move_tunnels() *LinkedRoomsList {
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
}*/

/*func (farm *Farm) Distribute_ant_starter() {
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
		// loop through each ant
		for ant_idx := range farm.ants {
			// check if this tunnel distance steps is lower than the next
			if sorted_first_tunnels_steps[sorted_head.room] < sorted_first_tunnels_steps[sorted_head.next.room] && !farm.ants[ant_idx].self_start {
				// Force the ant to move
				farm.ants[ant_idx].self_start = true
				farm.ants[ant_idx].force_move_to_room = sorted_head.room
				sorted_first_tunnels_steps[sorted_head.room]++
			}
		}

		sorted_head = sorted_head.next
	}
}*/

func (farm *Farm) same_distance_tunnels(ant *Ant) bool {
	tunnel := ant.room.tunnels.head
	similar_distance_count := 0

	for tunnel != nil {
		compare_tunnel := tunnel.next
		for compare_tunnel != nil {
			if farm.distances[tunnel.room] == farm.distances[compare_tunnel.room] {
				similar_distance_count++
			}
			compare_tunnel = compare_tunnel.next
		}
		tunnel = tunnel.next
	}

	return similar_distance_count > 1
}
