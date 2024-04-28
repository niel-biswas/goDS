package queue

type Queue struct {
	Values []string
}

func NewQueue() *Queue {
	return &Queue{
		Values: make([]string, 0),
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Values) == 0
}

func (q Queue) Peek() string {
	if q.IsEmpty() {
		return ""
	}
	return q.Values[0]
}

func (q *Queue) Enqueue(value string) {
	q.Values = append(q.Values, value)
}

func (q *Queue) EnqueueMultiple(values []string) {
	q.Values = append(q.Values, values...)
}

func (q *Queue) Dequeue() string {
	element := q.Values[0] // The first element is the one to be dequeued following the FIFO pattern.
	// if there is only 1 element in the queue then return the element and update slice to nil
	if len(q.Values) == 1 {
		q.Values = []string{}
		return element
	}
	q.Values = q.Values[1:] // slice off the element to be dequeued (similar to pop left stack operation).
	return element
}
