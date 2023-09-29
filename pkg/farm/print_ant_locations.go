package farm

import "fmt"

/*
Prints the location of each ant
*/
func (farm *Farm) Print_Ants_Locations() string {
	locations_as_string := ""

	for i, ant := range farm.ants {
		if ant.moved {
			//fmt.Printf("L%d-%s ", i+1, ant.room.name)
			locations_as_string += fmt.Sprintf("L%d-%s ", i+1, ant.room.name)
			ant.moved = false
		}
	}

	return locations_as_string + "\n"
}
