package main

type Node struct {
	Data interface{}
	Next *Node
}

// LIFO
type Stack struct {
	Top *Node
}

func (s *Stack) Push(data interface{}) {
	newNode := &Node{Data: data}

	if s.Top == nil {
		s.Top = newNode
	} else {
		newNode.Next = s.Top
		s.Top = newNode
	}
}

func (s *Stack) Pop() {
	if s.Top == nil {
		return
	} else if s.Top.Next == nil {
		s.Top = nil
	} else {
		s.Top = s.Top.Next
	}
}
