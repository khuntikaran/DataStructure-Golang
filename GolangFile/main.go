package main

import (
	"fmt"
)

func main() {

	la := Linkedlist{}
	la.add(5)
	la.add(6)
	la.add(3)
	la.add(7)
	la.add(2)
	la.add(8)
	la.display()
	la.addAt(9, 2)
	fmt.Printf("<<<<<<<<<after adding node at specific position>>>>>>>>\n")
	la.display()
	la.deleteAt(2)
	fmt.Printf("\n<<<<<<<<After deleting>>>>>>>\n ")
	la.display()
	la.search(7)
}

// here we implement singly and doubly linked list together
type Node struct {
	data int
	next *Node
	prev *Node
}

type Linkedlist struct {
	len  int
	head *Node
}

func (l *Linkedlist) add(value int) {
	n := &Node{data: value}

	if l.len == 0 {
		l.head = n
		l.len++
		return

	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = n
			n.prev = ptr
			l.len++
			return

		}
		ptr = ptr.next

	}
}

func (l *Linkedlist) display() {
	if l.len == 0 {
		fmt.Printf("Linked list is empty")

	}
	ptr := l.head
	for ptr.next != nil {
		ptr = ptr.next
	}
	if ptr.next == nil {
		for i := 0; i < l.len; i++ {
			fmt.Printf(" %v -> ", ptr.data)
			ptr = ptr.prev
		}
	}
}

func (l *Linkedlist) search(val int) {
	ptr := l.head
	for ptr.next != nil {
		if ptr.data == val {
			fmt.Printf("\n%v find it\n", val)

		}
		ptr = ptr.next
	}

}

func (l *Linkedlist) getpos(pos int) *Node {
	ptr := l.head
	for i := 0; i < pos; i++ {
		ptr = ptr.next

	}
	return ptr
}

func (l *Linkedlist) addAt(val int, pos int) {
	n := l.getpos(pos)
	p := &Node{data: val}
	q := l.getpos(pos - 1)
	q.next = p
	p.prev = q
	p.next = n
	n.prev = p
	l.len++
}

func (l *Linkedlist) deleteAt(pos int) {
	p := l.getpos(pos)
	q := l.getpos(pos - 1)
	//n := l.getpos(pos + 1)
	q.next = p.next
	p.next.prev = q
	l.len--
}
