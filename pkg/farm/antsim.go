package farm

import "fmt"

func (farm *Farm) AntSim_Step() {

	ants := farm.ants

	/* Loop throgh each ant */
	for ant_idx := range ants {
		tunnels := ants[ant_idx].room.short_tunnels.head

		/* Loop through each available tunnel */
		for tunnels != nil {
			/* If the tunneled room is empty, then move this ant to the next room */
			if (tunnels.room.is_empty || tunnels.room.end) && (!ants[ant_idx].discovered_rooms[tunnels.room] && ants[ant_idx].moving) {
				ants[ant_idx].discovered_rooms[tunnels.room] = true
				ants[ant_idx].room.is_empty = true // flag current room as empty
				ants[ant_idx].room = tunnels.room  // go to next room
				tunnels.room.is_empty = false      // flag next room as not empty
				if ants[ant_idx].room.end {
					ants[ant_idx].moving = false
				}
				break
			}
			tunnels = tunnels.next
		}
		farm.PrintShortFarm()
		farm.ReInitShortTunnels()
		farm.AntBFS()
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
