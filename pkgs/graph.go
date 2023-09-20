package pkgs

type EdgeNode struct {
	name      string    // Name of the edge
	x         int       // pos x
	y         int       // pos y
	tunnel_to string    // where the edge is connected
	next      *EdgeNode // The next edge in the list
}

type Graph struct {
	edges map[string]*EdgeNode
}
