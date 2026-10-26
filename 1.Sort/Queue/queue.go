package queue

type Queuer interface {
	Enqueue(v any)
	Dequeue() (any, bool)
	Peek() (any, bool)
	IsEmpty() bool
}

type node struct {
	v    any
	next *node
}

type Queue struct {
	head *node
	tail *node
}

func (q *Queue) Enqueue(v any) {
	newNode := &node{v: v}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) Dequeue() (any, bool) {
	if q.head == nil {
		return nil, false
	}
	removed := q.head.v
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return removed, true
}

func (q *Queue) Peek() (any, bool) {
	if q.head == nil {
		return nil, false
	}
	return q.head.v, true
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}
