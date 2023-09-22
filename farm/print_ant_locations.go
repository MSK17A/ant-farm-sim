package farm

import "fmt"

/*
	Prints the location of each ant
*/
func (farm *Farm) Print_Ants_Locations() {
	for i, ant_room := range farm.ants_rooms {
		fmt.Printf("ant %d in %s\n", i+1, ant_room.name)
	}
}
