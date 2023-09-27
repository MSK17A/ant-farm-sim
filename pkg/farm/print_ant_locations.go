package farm

import "fmt"

/*
Prints the location of each ant
*/
func (farm *Farm) Print_Ants_Locations() {
	for i, ant := range farm.ants {
		if ant.moved {
			fmt.Printf("L%d-%s ", i+1, ant.room.name)
			ant.moved = false
		}
	}
	fmt.Println()
}
