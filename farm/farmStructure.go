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
	Rooms map[string]*Room
	Ants  []Room
}
