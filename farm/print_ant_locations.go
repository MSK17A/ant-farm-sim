package farm

import "fmt"

/*
	Prints the location of each ant
*/
func (farm *Farm) Print_Ants_Locations() {
	for i, ant := range farm.ants {
		fmt.Printf("ant %d in %s\n", i+1, ant.room.name)
	}
}
