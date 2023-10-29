package main

import (
	ants "ants/pkg/farm"
	load_data "ants/pkg/farm_loading"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("USAGE: go run ./cmd <path_to_farm_file>")
		return
	}
	var farm ants.Farm
	farm.InitFarm()
	if !load_data.Read_Farm_File(args[0], &farm) {
		return
	}
	if !farm.Unique_Positions() {
		fmt.Println("Similar positions detected!")
		return
	}
	farm.InitDistances()
	farm.AntBFS()
	farm.PrintFarm()
	farm.PrintDistances()
	farm.AntSim()
}
