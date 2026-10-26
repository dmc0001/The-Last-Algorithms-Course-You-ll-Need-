package main

import (
	"fmt"
	"queue"
	"trees"
)

func breadthFirstSearch(tree *trees.Tree, target any) bool {
	if tree == nil {
		return false
	}
	q := queue.Queue{}
	q.Enqueue(tree)
	for !q.IsEmpty() {
		n, _ := q.Dequeue()
		curr := n.(*trees.Tree)
		if curr.Val == target {
			return true
		}
		if curr.Left != nil {
			q.Enqueue(curr.Left)
		}
		if curr.Right != nil {
			q.Enqueue(curr.Right)
		}
	}

	return false
}

func main() {
	//        1
	//      /   \
	//     2    3
	//    / \   /
	//   4   5 6
	root := trees.New(1)
	root.Left = trees.New(2)
	root.Left.Left = trees.New(4)
	root.Left.Right = trees.New(5)
	root.Right = trees.New(3)
	root.Right.Left = trees.New(6)
	fmt.Println("============================")
	fmt.Println(breadthFirstSearch(root, 5))  // true
	fmt.Println(breadthFirstSearch(root, 6))  // true
	fmt.Println(breadthFirstSearch(root, 99)) // false
}
