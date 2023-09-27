package main

import (
	ants "ants/pkg/farm"
)

func main() {
	var farm ants.Farm

	namla4(&farm)
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

func namla4(farm *ants.Farm) {
	farm.InitFarm()
	farm.AddRoom("richard", "start", 0, 0)
	farm.AddRoom("gilfoyle", "normal", 0, 0)
	farm.AddRoom("erlich", "normal", 0, 0)
	farm.AddRoom("dinish", "normal", 0, 0)
	farm.AddRoom("jimYoung", "normal", 0, 0)
	farm.AddRoom("peter", "end", 0, 0)
	farm.InitTunnels()
	farm.InitAnts(9)
	farm.InitDistances()

	farm.AddTunnel("richard", "dinish", true)
	farm.AddTunnel("dinish", "jimYoung", true)
	farm.AddTunnel("richard", "gilfoyle", true)
	farm.AddTunnel("gilfoyle", "peter", true)
	farm.AddTunnel("gilfoyle", "erlich", true)
	farm.AddTunnel("richard", "erlich", true)
	farm.AddTunnel("jimYoung", "peter", true)
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
	farm.AddRoom("0", "start", 0, 0)
	farm.AddRoom("4", "normal", 0, 0)
	farm.AddRoom("5", "end", 0, 0)
	farm.AddRoom("2", "normal", 0, 0)
	farm.AddRoom("3", "normal", 0, 0)
	farm.AddRoom("1", "normal", 0, 0)
	farm.InitTunnels()
	farm.InitAnts(4)
	farm.InitDistances()

	farm.AddTunnel("0", "1", true)
	farm.AddTunnel("2", "4", true)
	farm.AddTunnel("1", "4", true)
	farm.AddTunnel("0", "2", true)
	farm.AddTunnel("4", "5", true)
	farm.AddTunnel("3", "0", true)
	farm.AddTunnel("4", "3", true)
}

func namla1(farm *ants.Farm) {
	farm.InitFarm()
	farm.AddRoom("2", "normal", 0, 0)
	farm.AddRoom("3", "normal", 0, 0)
	farm.AddRoom("6", "normal", 0, 0)
	farm.AddRoom("7", "normal", 0, 0)
	farm.AddRoom("4", "normal", 0, 0)
	farm.AddRoom("5", "normal", 0, 0)
	farm.AddRoom("1", "start", 0, 0)
	farm.AddRoom("0", "end", 0, 0)
	farm.InitTunnels()
	farm.InitAnts(3)
	farm.InitDistances()

	farm.AddTunnel("0", "4", true)
	farm.AddTunnel("0", "6", true)
	farm.AddTunnel("1", "3", true)
	farm.AddTunnel("4", "3", true)
	farm.AddTunnel("5", "2", true)
	farm.AddTunnel("3", "5", true)
	farm.AddTunnel("4", "2", true)
	farm.AddTunnel("2", "1", true)
	farm.AddTunnel("7", "6", true)
	farm.AddTunnel("7", "2", true)
	farm.AddTunnel("7", "4", true)
	farm.AddTunnel("6", "5", true)
}
