package farm

type Room struct {
	tunnels        *LinkedRoomsList
	start          bool
	end            bool
	is_empty       bool
	pos_x          int
	pos_y          int
	name           string
	locked_tunnels map[string]bool
}

type Ant struct {
	room             *Room
	discovered_rooms map[*Room]bool
	moving           bool
}

type Farm struct {
	rooms          map[string]*Room
	number_of_ants int
	distances      map[*Room]int
	ants           []*Ant
	start_room     *Room
	end_room       *Room
}
