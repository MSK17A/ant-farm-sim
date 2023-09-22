package main

import ants "ants/farm"

func main() {
	var farm ants.Farm

	farm.InitFarm()
	farm.AddRoom("Room1", "start", 0, 0)
	farm.AddRoom("Room2", "normal", 0, 0)
	farm.AddRoom("Room3", "normal", 0, 0)
	farm.AddRoom("Room4", "normal", 0, 0)
	farm.AddRoom("Room5", "end", 0, 0)
	farm.InitAnts(3)

	farm.AddTunnel("Room1", "Room2")
	farm.AddTunnel("Room1", "Room4")
	farm.AddTunnel("Room2", "Room3")

	farm.PrintFarm()

}
