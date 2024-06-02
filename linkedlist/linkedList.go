package main

import (
	"fmt"
)

// If an identifier starts with a capital letter, it is exported.
// Meaning it can be accessed from other packages.

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(data interface{}) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head = newNode
	} else {
		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}

func (l *LinkedList) Print() {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Next != nil {
			fmt.Printf("%v -> ", currentNode.Data)
		} else {
			fmt.Printf("%v", currentNode.Data)
		}
		currentNode = currentNode.Next
	}
	fmt.Println()
}

// The iterative way to search linked list:
/*
func (l *LinkedList) Search(data interface{}) *Node {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Data == data {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}
*/

// Method can be called on a nil value, but
// accessing a field of a nil valued receiver results in a runtime panic
func (n *Node) Search(data interface{}) *Node {
	if n == nil {
		return nil
	}
	if n.Data == data {
		return n
	}
	return n.Next.Search(data)
}

func (l *LinkedList) Search(data interface{}) *Node {
	// we can access l.Head bc l can never be nil, but it can have nil head
	return l.Head.Search(data)
}

func (l *LinkedList) SearchAll(data interface{}) []*Node {
	var nodes []*Node
	for n := l.Head; n != nil; n = n.Next {
		if n.Data == data {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

func (l *LinkedList) Delete(data interface{}) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data {
		l.Head = l.Head.Next
		return
	}
	for n := l.Head; n != nil && n.Next != nil; n = n.Next {
		if n.Next.Data == data {
			n.Next = n.Next.Next
			return
		}
	}
}

func (l *LinkedList) DeleteAll(data interface{}) {
	for l.Head != nil && l.Head.Data == data {
		l.Head = l.Head.Next
	}
	if l.Head != nil {
		for n := l.Head; n.Next != nil; {
			if n.Next.Data == data {
				n.Next = n.Next.Next
			} else {
				n = n.Next
			}
		}
	}
}

func main() {
	list := &LinkedList{}
	list.Add(0)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(4)
	list.Add(4)
	list.Add(4)
	list.Add(5)
	//list.Search(4)
	//foundNode := list.SearchAll(4)
	//fmt.Println(foundNode)
	list.Print()
}
