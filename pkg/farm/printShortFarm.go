package farm

import "fmt"

/* Prints the farm and rooms connections */
func (farm *Farm) PrintShortFarm() {
	fmt.Println("Rooms connections:")

	for key, room := range farm.rooms {
		fmt.Printf("%s ->", key)
		short_tunnel := room.short_tunnels.head
		for short_tunnel != nil {
			fmt.Printf(" + %s", short_tunnel.room.name)
			short_tunnel = short_tunnel.next
		}
		fmt.Print("\n")
	}
}
