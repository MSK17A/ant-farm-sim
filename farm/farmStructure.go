package farm

type Room struct {
	tunnels  []*Room
	start    bool
	end      bool
	is_empty bool
	pos_x    int
	pos_y    int
	name     string
}

type Farm struct {
	rooms          map[string]*Room
	number_of_ants int
	ants_rooms     []*Room
	start_room     *Room
	end_room       *Room
}
