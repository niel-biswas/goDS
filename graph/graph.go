package graph

type Graph struct {
	AdjacencyList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		AdjacencyList: make(map[string][]string),
	}
}

func (g *Graph) AddNode(key string, val []string) {
	g.AdjacencyList[key] = val
}

func (g *Graph) hasCycleUtil(node string, visited map[string]bool, recStack map[string]bool, cycleNodes *[]string) bool {
	// The visited map keeps track of whether a node has been visited ever in the traversal
	visited[node] = true

	// The recStack keeps track of the nodes in the current path of the DFS traversal, i.e., the nodes that are still being processed
	// and have not been completely visited yet.
	recStack[node] = true

	for _, adjacentNode := range g.AdjacencyList[node] {
		if !visited[adjacentNode] {
			if g.hasCycleUtil(adjacentNode, visited, recStack, cycleNodes) {
				*cycleNodes = append(*cycleNodes, adjacentNode)
				return true
			}
		} else if recStack[adjacentNode] {
			*cycleNodes = append(*cycleNodes, adjacentNode)
			return true
		}
	}

	// If it has checked all adjacent nodes and no cycles are found, mark the node as not cyclic in the recursion stack
	recStack[node] = false
	return false
}

func (g *Graph) HasCycle() (bool, []string) {
	// This keeps track of all nodes that we have visited at least once during the entire execution of the algorithm.
	// Once a node is marked as visited, it remains marked for the rest of the algorithm. This is used to ensure that we don't process the same node more than once,
	// which could lead to an infinite loop in the case of a cyclic graph.
	visited := make(map[string]bool)

	// This keeps track of the nodes in the current path of the Depth-First Search (DFS). If we encounter a node that is already in the recStack,
	// it means we've come back to a node in the current path, i.e., a cycle exists.
	recStack := make(map[string]bool)

	// To add nodes that are cyclic
	cycleNodes := make([]string, 0)

	for i := range g.AdjacencyList {
		if !visited[i] {
			if g.hasCycleUtil(i, visited, recStack, &cycleNodes) {
				return true, cycleNodes
			}
		}
	}

	// If we used only the visited map, we wouldn't be able to detect cycles because we wouldn't have a way to distinguish between a node that has been visited
	// in the current path (which indicates a cycle) and a node that was visited in a previous path. The recStack gives us a way to keep track of the current path,
	// and that's why we need both the visited map and the recStack map for this algorithm.

	// Finally return the result
	return false, cycleNodes
}
