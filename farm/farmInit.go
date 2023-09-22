package farm

/* Allocates an initial space for the farm */
func (farm *Farm) InitFarm() {
	farm.rooms = make(map[string]*Room)
}

/* Puts all the ants in the start room */
func (farm *Farm) InitAnts(ants_number int) {
	farm.number_of_ants = ants_number
	farm.ants_rooms = make([]*Room, farm.number_of_ants)

	for i := 0; i < ants_number; i++ {
		farm.ants_rooms[i] = farm.start_room
	}
}