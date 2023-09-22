package pkgs

type Room struct {
	tunnels  []Room
	start    bool
	end      bool
	is_empty bool
	name     string
	pos_x    int
	pos_y    int
}

var ants []Room
