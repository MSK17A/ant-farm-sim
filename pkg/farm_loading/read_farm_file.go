package farmloading

import (
	"ants/pkg/farm"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func Read_Farm_File(file_path string, farm *farm.Farm) {
	file, err := os.Open(file_path)
	check(err)

	defer file.Close()
	line_number := 0

	scanner := bufio.NewScanner(file)
	start_flag := false
	end_flag := false
	n_of_ants := 0

	for scanner.Scan() {
		line_number++
		if line_number == 1 {
			n_of_ants_line_fields := strings.Fields(scanner.Text())
			if len(n_of_ants_line_fields) != 1 {
				fmt.Println("Error in reading number of ants")
				os.Exit(1)
			}
			n_of_ants, _ = strconv.Atoi(n_of_ants_line_fields[0])
		}
		if start_flag && end_flag {
			fmt.Println("No start or end room")
			os.Exit(1)
		}
		if scanner.Text() == "##start" {
			start_flag = true
			continue
		}
		if scanner.Text() == "##end" {
			end_flag = true
			continue
		}
		if start_flag {
			start_flag = false
			args := strings.Fields(scanner.Text())
			if len(args) != 3 {
				fmt.Printf("Error in line %d\n", line_number)
				continue
			}
			name := args[0]
			pos_x, _ := strconv.Atoi(args[1])
			pos_y, _ := strconv.Atoi(args[2])
			farm.AddRoom(name, "start", pos_x, pos_y)
			continue
		}
		if end_flag {
			end_flag = false
			args := strings.Fields(scanner.Text())
			if len(args) != 3 {
				fmt.Printf("Error in line %d\n", line_number)
				continue
			}
			name := args[0]
			pos_x, _ := strconv.Atoi(args[1])
			pos_y, _ := strconv.Atoi(args[2])
			farm.AddRoom(name, "end", pos_x, pos_y)
			continue
		}

		if strings.Contains(scanner.Text(), "-") && strings.Count(scanner.Text(), "-") == 1 {
			args := strings.Split(scanner.Text(), "-")
			farm.AddTunnel(args[0], args[1], true)
			continue
		}

		room_name_loc := strings.Fields(scanner.Text())
		if len(room_name_loc) != 3 && (!start_flag || !end_flag) {
			//fmt.Printf("Error in line %d\n", line_number)
			//It might hit a comment
			continue
		}
		name := room_name_loc[0]
		pos_x, _ := strconv.Atoi(room_name_loc[1])
		pos_y, _ := strconv.Atoi(room_name_loc[2])
		farm.AddRoom(name, "normal", pos_x, pos_y)
	}
	if n_of_ants == 0 {
		fmt.Println("No Ants!!!")
	}
	farm.InitAnts(n_of_ants)
}
