package main

import (
	"trees"
)

func depthFirstSearch(tree *trees.Tree, target int) bool {
	if tree == nil {
		return false
	}
	if tree.Val == target {
		return true
	}
	if tree.Val <= target {
		return depthFirstSearch(tree.Left, target)
	}
	if tree.Val > target {
		return depthFirstSearch(tree.Right, target)
	}
	return false
}

func main() {
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6

	// Find
}
