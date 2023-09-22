package pkgs

import "fmt"

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

/* Allocates an initial space for the farm */
func (farm *Farm) InitFarm() {
	farm.Rooms = make(map[string]*Room)
}

/* Add a room to the farm */
func (farm *Farm) AddRoom(name string, start_end_normal string, pos_x int, pos_y int) {
	var new_room Room

	switch start_end_normal {
	case "start":
		farm.Rooms[name] = &Room{
			name:     name,
			start:    true,
			end:      false,
			is_empty: true,
			pos_x:    pos_x,
			pos_y:    pos_y,
		}
	case "end":
		new_room = Room{
			name:     name,
			start:    false,
			end:      true,
			is_empty: true,
			pos_x:    pos_x,
			pos_y:    pos_y,
		}
		farm.Rooms[name] = &new_room
	case "normal":
		new_room = Room{
			name:     name,
			start:    false,
			end:      false,
			is_empty: true,
			pos_x:    pos_x,
			pos_y:    pos_y,
		}
		farm.Rooms[name] = &new_room
	default:
		fmt.Println("Enter room type")
	}
}

/* Joins two rooms together */
func (farm *Farm) AddTunnel(from_room string, to_room string) {
	farm.Rooms[from_room].tunnels = append(farm.Rooms[from_room].tunnels, farm.Rooms[to_room])
}

/* Prints the farm and rooms connections */
func (farm *Farm) PrintFarm() {
	for key, room := range farm.Rooms {
		fmt.Printf("%s", key)
		if len(room.tunnels) > 0 {
			for _, tunnel := range room.tunnels {
				fmt.Printf(" -> %s", tunnel.name)
			}
		}
		fmt.Print("\n")
	}
}
