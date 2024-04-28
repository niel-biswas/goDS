package stack

type Stack struct {
	Values []string
}

func NewStack() *Stack {
	return &Stack{
		Values: make([]string, 0),
	}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(value string) {
	s.Values = append(s.Values, value)
}

// Pop removes an item from the top of the stack
func (s *Stack) Pop() string {
	length := len(s.Values)
	// Extract the top item
	toRemove := s.Values[length-1]
	// Remove the top item
	s.Values = s.Values[:length-1]
	return toRemove
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.Values) == 0
}

// Peek returns the item from the top of the stack without removing it
func (s *Stack) Peek() string {
	length := len(s.Values)
	return s.Values[length-1]
}
