package main

import "fmt"

func main() {
	var t Tree
	t.Insert(500)
	t.Insert(300)
	t.Insert(100)
	t.Insert(600)
	t.Insert(200)
	t.Insert(1000)
	t.Insert(700)
	t.Insert(345)
	t.Insert(230)
	t.Insert(555)
	t.Insert(666)
	//Inoreder(t.root)
	//fmt.Printf("\n")
	//Postoreder(t.root)
	//fmt.Printf("\n")
	//Preoreder(t.root)
	t.Search(345)
	fmt.Printf("\n")
	t.minval()
	fmt.Printf("\n")
	t.maxval()

}

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &Node{data: val}
	} else {
		t.root.Insert(val)
	}
}

func (n *Node) Insert(val int) {
	if val <= n.data {
		if n.left == nil {
			n.left = &Node{data: val}
		} else {
			n.left.Insert(val)
		}

	} else {
		if n.right == nil {
			n.right = &Node{data: val}

		} else {
			n.right.Insert(val)
		}
	}

}

func Inoreder(n *Node) {
	if n == nil {
		return
	} else {

		Inoreder(n.left)
		fmt.Printf("%d-> ", n.data)
		Inoreder(n.right)
	}
}
func Postoreder(n *Node) {
	if n == nil {
		return
	} else {

		Postoreder(n.left)
		Postoreder(n.right)
		fmt.Printf("%d-> ", n.data)

	}
}

func Preoreder(n *Node) {
	if n == nil {
		return
	} else {

		Preoreder(n.left)
		Preoreder(n.right)
		fmt.Printf("%d-> ", n.data)

	}
}

func (t *Tree) Search(val int) {
	if val == t.root.data {
		fmt.Printf("Found %d ", t.root.data)
	} else {
		t.root.Search(val)
	}

}

func (n *Node) Search(val int) {
	if val == n.data {
		fmt.Printf(" found it %d", n.data)
		return
	}
	if val < n.data {
		if val <= n.left.data {

			n.left.Search(val)
			return

		}
		if val > n.left.data {
			n.left.Search(val)
			return

		}

	} else {
		if val <= n.right.data {
			n.right.Search(val)
			return

		}
		if val > n.right.data {
			n.right.Search(val)
			return

		}
	}
}

func (t *Tree) minval() *Node {
	for t.root.left != nil {
		t.root = t.root.left

	}
	fmt.Printf("%d", t.root.data)
	return t.root

}

func (t *Tree) maxval() *Node {
	for t.root.right != nil {
		t.root = t.root.right
	}
	fmt.Printf("%d", t.root.data)
	return t.root
}
