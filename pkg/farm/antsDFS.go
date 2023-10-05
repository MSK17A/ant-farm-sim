package farm

import "fmt"

var (
	visited = make(map[string]bool)
)

/* Depth First Search with marking dead end tunnels */
func (farm *Farm) AntDFS(tunnel_head *RoomNode) bool {
	if tunnel_head == nil {
		return false // Handle the case where tunnel_head is nil
	}

	visited[farm.start_room.name] = true

	if !visited[tunnel_head.room.name] || tunnel_head.room.end {
		visited[tunnel_head.room.name] = true
		fmt.Printf("Visited %s\n", tunnel_head.room.name)
		if tunnel_head.room.end {
			return false // We've reached the end, not a dead end
		}
		if !farm.AntDFS(tunnel_head.room.tunnels.head) {
			// If the subpath is a dead end, mark the current room as a dead end.
			tunnel_head.room.dead_end = true
		}
	}
	if farm.AntDFS(tunnel_head.next) {
		// If the next subpath is a dead end, mark the current room as a dead end.
		tunnel_head.room.dead_end = true
	}
	return !tunnel_head.room.dead_end
}
