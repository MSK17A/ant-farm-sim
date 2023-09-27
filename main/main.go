package main

import (
	ants "ants/pkg/farm"
)

func main() {
	var farm ants.Farm

	namla2(&farm)
	/* Print farm */
	//farm.PrintFarm()
	farm.AntBFS()
	//farm.PrintDistances()

	farm.AntSim()
	/*
		fmt.Println("\nStep 5")
		farm.AntSim_Step()
		farm.Print_Ants_Locations()
	*/
}

func namla3(farm *ants.Farm) {
	farm.InitFarm()
	farm.AddRoom("0", "start", 0, 0)
	farm.AddRoom("1", "normal", 0, 0)
	farm.AddRoom("2", "normal", 0, 0)
	farm.AddRoom("3", "end", 0, 0)
	farm.InitTunnels()
	farm.InitAnts(20)
	farm.InitDistances()

	farm.AddTunnel("0", "1", true)
	farm.AddTunnel("0", "3", true)
	farm.AddTunnel("1", "2", true)
	farm.AddTunnel("3", "2", true)
}

func namla2(farm *ants.Farm) {

	farm.InitFarm()
	farm.AddRoom("Room0", "start", 0, 0)
	farm.AddRoom("Room4", "normal", 0, 0)
	farm.AddRoom("Room5", "end", 0, 0)
	farm.AddRoom("Room2", "normal", 0, 0)
	farm.AddRoom("Room3", "normal", 0, 0)
	farm.AddRoom("Room1", "normal", 0, 0)
	farm.InitTunnels()
	farm.InitAnts(4)
	farm.InitDistances()

	farm.AddTunnel("Room0", "Room1", true)
	farm.AddTunnel("Room2", "Room4", true)
	farm.AddTunnel("Room1", "Room4", true)
	farm.AddTunnel("Room0", "Room2", true)
	farm.AddTunnel("Room4", "Room5", true)
	farm.AddTunnel("Room3", "Room0", true)
	farm.AddTunnel("Room4", "Room3", true)
}

func namla1(farm *ants.Farm) {
	farm.InitFarm()
	farm.AddRoom("Room2", "normal", 0, 0)
	farm.AddRoom("Room3", "normal", 0, 0)
	farm.AddRoom("Room6", "normal", 0, 0)
	farm.AddRoom("Room7", "normal", 0, 0)
	farm.AddRoom("Room4", "normal", 0, 0)
	farm.AddRoom("Room5", "normal", 0, 0)
	farm.AddRoom("Room1", "start", 0, 0)
	farm.AddRoom("Room0", "end", 0, 0)
	farm.InitTunnels()
	farm.InitAnts(3)
	farm.InitDistances()

	farm.AddTunnel("Room0", "Room4", true)
	farm.AddTunnel("Room0", "Room6", true)
	farm.AddTunnel("Room1", "Room3", true)
	farm.AddTunnel("Room4", "Room3", true)
	farm.AddTunnel("Room5", "Room2", true)
	farm.AddTunnel("Room3", "Room5", true)
	farm.AddTunnel("Room4", "Room2", true)
	farm.AddTunnel("Room2", "Room1", true)
	farm.AddTunnel("Room7", "Room6", true)
	farm.AddTunnel("Room7", "Room2", true)
	farm.AddTunnel("Room7", "Room4", true)
	farm.AddTunnel("Room6", "Room5", true)
}
