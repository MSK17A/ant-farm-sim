package pkgs

import "fmt"

/* Store the edge between one node to node */
type EdgeNode struct {
	name string // Name of the edge
	//x         int       // pos x
	//y         int       // pos y
	Connected_to string    // where the edge is connected
	Next         *EdgeNode // The Next edge in the list
}

/* Stores the edges in a map */
type Graph struct {
	Edges map[string]*EdgeNode
}

func (graph *Graph) InitGraph() {
	graph.Edges = make(map[string]*EdgeNode)
}

/* This function will add an edge from one node to another, it will add the node if it is not available */
func (graph *Graph) AddNode(name string, connected_to string) {

	if graph.Edges[name] == nil {
		graph.Edges[name] = &EdgeNode{
			name:         name,
			Connected_to: connected_to,
			Next:         nil,
		}
	} else {
		edgenode := lastEdgeNode(graph.Edges[name], connected_to)
		if edgenode == nil {
			return
		}
		edgenode.Next = &EdgeNode{
			name:         name,
			Connected_to: connected_to,
			Next:         nil,
		}
	}
}

/* Represents the graph nodes connections */
func (graph *Graph) PrintGraph() {

	for _, value := range graph.Edges {
		//fmt.Printf("%s -> %s\n", key, value.Connected_to)
		PrintAdj(value)
	}
}

/* This function will traverse through all edges the node connected to
 */
func lastEdgeNode(edgenode *EdgeNode, connected_to string) *EdgeNode {
	if edgenode.Connected_to == connected_to {
		return nil
	}
	if edgenode.Next == nil {
		return edgenode
	}
	return lastEdgeNode(edgenode.Next, connected_to)
}

/* Prints all the edges the node connected to */
func PrintAdj(edgenode *EdgeNode) {
	if edgenode == nil {
		return
	}

	fmt.Printf("%s -> %s\n", edgenode.name, edgenode.Connected_to)
	PrintAdj(edgenode.Next)
}
