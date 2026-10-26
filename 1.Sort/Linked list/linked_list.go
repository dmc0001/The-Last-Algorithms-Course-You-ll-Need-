package main

import "fmt"

func main() {
	root := LinkedList()
	root.insert(10)
	root.insert(11)
	root.insert(12)
	root.insert(13)
	root.update(12, 100)
	root.insert(14)
	root.delete(14)
	root.insertAt(900, 2)
	root.read()
	fmt.Printf("The root length is: %d\n", root.length())

}

type Node struct {
	v    int
	next *Node
}

func LinkedList() *Node {
	return &Node{}
}

func (n *Node) insert(v int) {
	if n.next == nil {
		n.next = &Node{v: v}
		return
	}
	n.next.insert(v)
}

func (n *Node) insertAt(v int, p int) {
	if p <= 0 {
		newNode := &Node{v: v, next: n.next}
		n.next = newNode
		return
	}
	count := 0
	if n.next == nil {
		n.next = &Node{v: v}
		return
	}
	curr := n
	for curr != nil && count < p-1 {
		count++
		curr = curr.next
	}
	if curr != nil {
		newNode := &Node{v: v, next: curr.next}
		curr.next = newNode
	}

}

func (n *Node) delete(v int) {
	if n == nil {
		return
	}
	if n.next.v == v {
		n.next = n.next.next
		return
	}
	n.next.delete(v)

}

func (n *Node) update(v int, nv int) {
	if n.next == nil || n == nil {
		return
	}
	if n.next.v == v {
		n.next.v = nv
		return
	}
	n.next.update(v, nv)
}

func (n *Node) read() {
	for curr := n; curr != nil; curr = curr.next {
		fmt.Println(curr.v)
	}
}

func (n *Node) length() int {
	l := 0
	for curr := n; curr != nil; curr = curr.next {
		l++
	}
	return l
}
