package farm

func Number_of_Tunnels(tunnels *LinkedRoomsList) int {

	counter := 0

	tunnel := tunnels.head

	for tunnel != nil {
		counter++
		tunnel = tunnel.next
	}

	return counter
}
