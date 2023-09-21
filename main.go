package main

import ants "ants/pkgs"

func main() {
	var farm ants.Graph
	farm.InitGraph()

	farm.AddNode("room1", "room2")
	farm.AddNode("room1", "room3")
	farm.AddNode("room2", "room3")

	//ants.PrintAdj(farm.Edges["room1"])

	farm.PrintGraph()

}
