package farm

import "fmt"

/* Prints the farm and rooms connections */
func (farm *Farm) PrintFarm() {
	for key, room := range farm.rooms {
		fmt.Printf("%s ->", key)
		if len(room.tunnels) > 0 {
			for _, tunnel := range room.tunnels {
				fmt.Printf(" + %s", tunnel.name)
			}
		}
		fmt.Print("\n")
	}
}
