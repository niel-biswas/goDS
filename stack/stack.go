package stack

type Stack struct {
	Nodes []string
	Pi    string // pi represents the parent node
	Depth int
}

func NewStack() *Stack {
	return &Stack{
		Nodes: make([]string, 0),
	}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(values string, pi string, depth int) {
	s.Nodes = append(s.Nodes, values)
	s.Pi = pi
	s.Depth = depth
}

// Pop removes an item from the top of the stack
func (s *Stack) Pop() string {
	length := len(s.Nodes)
	// Extract the top item
	toRemove := s.Nodes[length-1]
	// Remove the top item
	s.Nodes = s.Nodes[:length-1]
	return toRemove
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.Nodes) == 0
}

// Peek returns the item from the top of the stack without removing it
func (s *Stack) Peek() string {
	length := len(s.Nodes)
	return s.Nodes[length-1]
}
