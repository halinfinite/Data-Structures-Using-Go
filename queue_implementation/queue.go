package main

type Node struct {
	Data interface{}
	Next *Node
}

// FIFO
type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enqueue(data interface{}) {
	newNode := &Node{Data: data}

	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
	} else if q.Head.Next == nil {
		q.Head.Next = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
}

func (q *Queue) Dequeue() {
	if q.Head == nil {
		return
	} else if q.Head.Next == nil {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Next
	}
}
