package farm

import "fmt"

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
