package farm

/*
Checks if ants are stuck and can't move
It will check even the ant that previosly checked and it replied with no tunnels,
pherhaps change the structure to queue or stacks? linked lists?
*/
func (farm *Farm) Ants_Stuck() bool {

	for _, ant := range farm.ants {
		if len(ant.room.tunnels) > 0 && !ant.room.end {
			for _, tunnel := range ant.room.tunnels {
				if !ant.discovered_rooms[tunnel] {
					return false
				}
			}
		}
	}
	return true
}
