package graph

import "fmt"

type WeightedGraph struct {
	Nodes         map[string]map[string]int // key is the node and value represents the edge b/w nodes having weights
	bidirectional bool                      // true if the graph is undirected, false for directed graphs
}

func NewWeightedGraph(bidirectional bool) *WeightedGraph {
	return &WeightedGraph{
		Nodes:         make(map[string]map[string]int),
		bidirectional: bidirectional,
	}
}

func (g *WeightedGraph) addNode(key string) {
	if _, exists := g.Nodes[key]; !exists {
		g.Nodes[key] = make(map[string]int)
	}
}

func (g *WeightedGraph) AddEdges(start, end string, weight int) {
	if weight < 0 {
		fmt.Println("Warning: Dijkstra's algorithm does not work correctly with negative weights. Found negative weight", weight, "from", start, "to", end)
	}
	g.addNode(start)
	g.addNode(end)
	g.Nodes[start][end] = weight
	if g.bidirectional {
		g.Nodes[end][start] = weight
	}
}

func (g *WeightedGraph) hasCycleUtil(v string, visited map[string]bool, recStack map[string]bool, cyclicNodes *[]string) bool {
	visited[v] = true
	recStack[v] = true

	for i := range g.Nodes[v] {
		if !visited[i] {
			if g.hasCycleUtil(i, visited, recStack, cyclicNodes) {
				*cyclicNodes = append(*cyclicNodes, i)
				return true
			}
		} else if recStack[i] {
			*cyclicNodes = append(*cyclicNodes, i)
			return true
		}
	}

	recStack[v] = false
	return false
}

func (g *WeightedGraph) HasCycle() (bool, []string) {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)
	cyclicNodes := make([]string, 0)

	for i := range g.Nodes {
		if !visited[i] {
			if g.hasCycleUtil(i, visited, recStack, &cyclicNodes) {
				return true, cyclicNodes
			}
		}
	}

	return false, cyclicNodes
}
