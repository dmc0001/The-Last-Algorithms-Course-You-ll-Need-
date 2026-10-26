package main

import "fmt"

type Node struct {
	v        any
	next     *Node
	previous *Node
}

// insert appends a new node with value nv at the tail of the list.
// If the head node is empty (zero value), it is initialised in place.
func (dl *Node) insert(nv int) {
	if dl.v == nil && dl.next == nil {
		dl.v = nv
		return
	}

	curr := dl

	// Move to last node
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &Node{v: nv, previous: curr}
}

// prepend inserts a new node with value nv before the current head.
// Because the caller's pointer must remain the head, the current head's
// content is pushed into a new second node and the head is overwritten with nv.
func (dl *Node) prepend(nv int) {
	oldHead := &Node{v: dl.v, next: dl.next, previous: dl}
	if dl.next != nil {
		dl.next.previous = oldHead
	}
	dl.v = nv
	dl.next = oldHead
}

// insertAt inserts a new node with value nv at zero-based index i.
// If i is 0 the new node becomes the new head; otherwise it is spliced
// in after the node currently at index i-1.
func (dl *Node) insertAt(nv int, i int) {
	if dl == nil {
		return
	}

	if i == 0 {
		dl.prepend(nv)
		return
	}

	curr := dl
	count := 0

	// Move to exact index
	for count < i-1 && curr.next != nil {
		curr = curr.next
		count++
	}

	next := curr.next
	curr.next = &Node{v: nv, previous: curr, next: next}
	if next != nil {
		next.previous = curr.next
	}
}

// delete removes the first node whose value equals v, relinking the
// surrounding nodes to preserve the doubly-linked structure.
func (dl *Node) delete(v int) {
	if dl == nil {
		return
	}

	curr := dl

	// handle head deletion
	if curr.v == v {
		if curr.next != nil {
			curr.next.previous = nil
			*dl = *curr.next
		} else {
			dl.v = nil
		}
		return
	}

	for curr != nil && curr.v != v {
		curr = curr.next
	}

	if curr == nil {
		return
	}

	if curr.previous != nil {
		curr.previous.next = curr.next
	}

	if curr.next != nil {
		curr.next.previous = curr.previous
	}
}

// get returns the value at zero-based index i and true if the index
// is in range, or nil and false if it is out of bounds.
func (dl *Node) get(i int) (any, bool) {
	curr := dl
	count := 0
	for curr != nil {
		if count == i {
			return curr.v, true
		}
		curr = curr.next
		count++
	}
	return nil, false
}

// contains reports whether any node in the list holds value v.
func (dl *Node) contains(v int) bool {
	curr := dl
	for curr != nil {
		if curr.v == v {
			return true
		}
		curr = curr.next
	}
	return false
}

// length returns the number of nodes in the list.
func (dl *Node) length() int {
	count := 0
	curr := dl
	for curr != nil && curr.v != nil {
		count++
		curr = curr.next
	}
	return count
}

// reverse reverses the list in place by swapping node values from both ends
// toward the middle, leaving the next/previous pointer structure unchanged.
func (dl *Node) reverse() {
	if dl == nil || dl.next == nil {
		return
	}

	left := dl
	right := dl
	for right.next != nil {
		right = right.next
	}

	for left != right && left.previous != right {
		left.v, right.v = right.v, left.v
		left = left.next
		right = right.previous
	}
}

// print traverses the list from head to tail, printing each value.
func (dl *Node) print() {
	curr := dl
	for curr != nil && curr.v != nil {
		fmt.Println(curr.v)
		curr = curr.next
	}
}

func main() {
	root := &Node{}
	root.insert(1)
	root.insert(12)
	root.insert(13)
	root.insert(14)

	fmt.Println("--- after inserts ---")
	root.print()

	root.prepend(0)
	fmt.Println("--- after prepend(0) ---")
	root.print()

	root.insertAt(100, 2)
	fmt.Println("--- after insertAt(100, 2) ---")
	root.print()

	root.delete(12)
	fmt.Println("--- after delete(12) ---")
	root.print()

	fmt.Println("length:", root.length())
	fmt.Println("contains 100:", root.contains(100))
	fmt.Println("contains 99:", root.contains(99))

	v, ok := root.get(2)
	fmt.Println("get(2):", v, ok)

	root.reverse()
	fmt.Println("--- after reverse ---")
	root.print()
}
