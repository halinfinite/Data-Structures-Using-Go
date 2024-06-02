package main

type Node struct {
	Data   int
	Parent *Node
	Left   *Node
	Right  *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) Insert(data int) {
	newNode := &Node{Data: data}
	var parent *Node // := would also assign a value to it, we want nil

	currentNode := bst.Root
	for currentNode != nil {
		parent = currentNode
		if data < currentNode.Data {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}

	newNode.Parent = parent

	if parent == nil {
		bst.Root = newNode // Given empty tree case
	} else if data < parent.Data {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
}

func (bst *BinarySearchTree) Search(n *Node, data int) *Node {
	if n == nil || n.Data == data {
		return n
	}
	if data < n.Data {
		return bst.Search(n.Left, data)
	} else {
		return bst.Search(n.Right, data)
	}
}

func (bst *BinarySearchTree) FindMin() *Node {
	currentNode := bst.Root
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}

func (n *Node) TreeMin() *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (bst *BinarySearchTree) FindMax() *Node {
	currentNode := bst.Root
	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}
	return currentNode
}

func (bst *BinarySearchTree) Transplant(u, v *Node) {
	if u.Parent == nil {
		bst.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

func (bst *BinarySearchTree) Delete(n *Node) {
	if n.Left == nil {
		bst.Transplant(n, n.Right)
	} else if n.Right == nil {
		bst.Transplant(n, n.Left)
	} else {
		y := n.Right.TreeMin()
		if y != n.Right {
			bst.Transplant(y, y.Right)
			y.Right = n.Right
			y.Right.Parent = y
		}
		bst.Transplant(n, y)
		y.Left = n.Left
		y.Left.Parent = y
	}
}
