package farm

/* Allocates an initial space for the farm */
func (farm *Farm) InitFarm() {
	farm.rooms = make(map[string]*Room)
}

/* Puts all the ants in the start room */
func (farm *Farm) InitAnts(ants_number int) {
	farm.number_of_ants = ants_number
	farm.ants = make([]*Ant, farm.number_of_ants)

	for i := 0; i < ants_number; i++ {
		farm.ants[i] = new(Ant)
		farm.ants[i].room = farm.start_room
		farm.ants[i].discovered_rooms = make(map[*Room]bool)
		farm.ants[i].discovered_rooms[farm.start_room] = true
		farm.ants[i].moving = true
	}
}
