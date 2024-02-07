package queue

type Queue struct {
	Nodes []string
	Pi    string // pi represents the parent node
	Depth int
	Seen  []string
}

func NewQueue() *Queue {
	return &Queue{
		Nodes: make([]string, 0),
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Nodes) == 0
}

func (q *Queue) Peek() string {
	if q.IsEmpty() {
		return ""
	}
	return q.Nodes[0]
}

func (q *Queue) IsSeen(node string) bool {
	for _, v := range q.Seen {
		if v == node {
			return true
		}
	}
	return false
}

func (q *Queue) Enqueue(values []string, pi string, depth int, seen string) {
	q.Nodes = append(q.Nodes, values...)
	q.Pi = pi
	q.Depth = depth
	q.Seen = append(q.Seen, seen)
}

func (q *Queue) Dequeue() string {
	element := q.Nodes[0] // The first element is the one to be dequeued following the FIFO pattern.
	// if there is only 1 element in the queue then return the element and update slice to nil
	if len(q.Nodes) == 1 {
		q.Nodes = []string{}
		return element
	}
	q.Nodes = q.Nodes[1:] // slice off the element to be dequeued (similar to pop left stack operation).
	return element
}
