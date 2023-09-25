package main

import (
	ants "ants/pkg/farm"
)

func main() {
	var farm ants.Farm

	farm.InitFarm()
	farm.AddRoom("Room2", "normal", 0, 0)
	farm.AddRoom("Room3", "normal", 0, 0)
	farm.AddRoom("Room6", "normal", 0, 0)
	farm.AddRoom("Room7", "normal", 0, 0)
	farm.AddRoom("Room4", "normal", 0, 0)
	farm.AddRoom("Room5", "normal", 0, 0)
	farm.AddRoom("Room1", "start", 0, 0)
	farm.AddRoom("Room0", "end", 0, 0)
	farm.InitAnts(3)

	farm.AddTunnel("Room0", "Room4", true)
	farm.AddTunnel("Room3", "Room5", true)
	farm.AddTunnel("Room1", "Room3", true)
	farm.AddTunnel("Room5", "Room2", true)
	farm.AddTunnel("Room7", "Room4", true)
	farm.AddTunnel("Room7", "Room2", true)
	farm.AddTunnel("Room0", "Room6", true)
	farm.AddTunnel("Room4", "Room2", true)
	farm.AddTunnel("Room2", "Room1", true)
	farm.AddTunnel("Room7", "Room6", true)
	farm.AddTunnel("Room4", "Room3", true)
	farm.AddTunnel("Room6", "Room5", true)

	/* Print farm */
	farm.PrintFarm()

	farm.AntBFS()
	farm.AntSim()
}
