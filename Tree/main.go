package main

import "fmt"

func main() {
	n := Node{}
	a := n.CreateNode(100)
	a.Insert(35)
	a.Insert(500)
	a.Insert(600)
	a.Insert(2)
	a.Insert(55)
	a.Insert(101)
	InorderTraverse(a)
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (n *Node) CreateNode(val int) *Node {
	root := &Node{data: val, left: nil, right: nil}
	return root
}

func (n *Node) Insert(val int) {

	if n == nil {
		n = &Node{data: val, left: nil, right: nil}

	} else {
		if n.left == nil && val <= n.data {
			n.left = n.CreateNode(val)
		}
		if val <= n.data && n.left != nil {
			if val <= n.left.data {
				n.left.left = n.CreateNode(val)
			} else {
				n.left.right = n.CreateNode(val)
			}

		}
		if val > n.data && n.right == nil {
			n.right = n.CreateNode(val)

		}
		if val > n.data && n.right != nil {
			if val > n.right.data {
				n.right.right = n.CreateNode(val)
			} else {
				n.right.left = n.CreateNode(val)
			}

		}

	}

}

func (n *Node) InsertLeft(root *Node, val int) *Node {
	root.left = n.CreateNode(val)
	return root.left
}

func (n *Node) InsertRight(root *Node, val int) *Node {
	root.right = n.CreateNode(val)
	return root.right
}

func InorderTraverse(t *Node) {
	if t == nil {
		return
	} else {
		fmt.Printf("%d-> ", t.data)
		InorderTraverse(t.left)

		InorderTraverse(t.right)

	}

}
