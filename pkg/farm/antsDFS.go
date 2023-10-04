package farm

var (
	visited    = make(map[string]bool)
	discovered = make(map[string]bool)
)

func (farm *Farm) AntDFS(tunnel_head *RoomNode) bool {
	if tunnel_head != nil {
		if !visited[tunnel_head.room.name] {
			visited[tunnel_head.room.name] = true
			tunnel_head.room.dead_end = farm.AntDFS(tunnel_head.room.tunnels.head.next)
		}
		if tunnel_head.room.end {
			return false
		}
	}
	tunnel_head.room.dead_end = farm.AntDFS(tunnel_head.next)
	return true
}
