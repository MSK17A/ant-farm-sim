package farm

/*
	Checks if ants are stuck and can't move
	It will check even the ant that previosly checked and it replied with no tunnels,
	pherhaps change the structure to queue or stacks? linked lists?
*/
func (farm *Farm) Ants_Stuck() bool {

	for _, ant_room := range farm.ants_rooms {
		if len(ant_room.tunnels) > 0 {
			return false
		}
	}
	return true
}
