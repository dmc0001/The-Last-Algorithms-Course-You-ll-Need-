package trees

type Tree struct {
	Left  *Tree
	Val   int
	Right *Tree
}

func New(val int) *Tree {
	return &Tree{Val: val}
}

func (tree *Tree) Insert(value int) {
	newNode := &Tree{Val: value}
	if tree == nil {
		tree = newNode
		return
	}
	if value < tree.Val {
		if tree.Left == nil {
			tree.Left = newNode
			return
		}
		tree.Left.Insert(value)

	}
	if value > tree.Val {
		if tree.Right == nil {
			tree.Right = newNode
			return
		}
		tree.Right.Insert(value)
	}

}

// Delete removes target and returns the (possibly new) subtree root.
// Call as: root = root.Delete(target)
func (tree *Tree) Delete(target int) *Tree {
	if tree == nil {
		return nil // not found
	}

	if target < tree.Val {
		tree.Left = tree.Left.Delete(target)
		return tree
	}
	if target > tree.Val {
		tree.Right = tree.Right.Delete(target)
		return tree
	}

	// Found it.
	if tree.Left == nil {
		return tree.Right // 0 or 1 (right) child
	}
	if tree.Right == nil {
		return tree.Left // 1 (left) child
	}

	// Two children: copy in-order successor's value, then delete it.
	succ := tree.Right.getSuccesor()
	tree.Val = succ.Val
	tree.Right = tree.Right.Delete(succ.Val)
	return tree
}

// depthFirstSearch
func (tree *Tree) depthFirstSearch(value int) *Tree {
	if tree == nil {
		return nil
	}
	if tree.Val == value {
		return tree
	}
	if tree.Val > value {
		return tree.Left.depthFirstSearch(value)
	}
	return tree.Right.depthFirstSearch(value)
}

// isLeaf
func (tree *Tree) isLeaf() bool {
	if tree.Left == nil && tree.Right == nil {
		return true
	}
	return false
}

// getParent
func (tree *Tree) getParent(target int) *Tree {

	if tree == nil {
		return nil
	}
	if (tree.Left != nil && tree.Left.Val == target) ||
		(tree.Right != nil && tree.Right.Val == target) {
		return tree
	}

	if tree.Val < target {
		return tree.Right.getParent(target)
	}

	return tree.Left.getParent(target)
}

// getSuccesor
func (tree *Tree) getSuccesor() *Tree {
	if tree == nil {
		return nil
	}
	if tree.Left == nil {
		return tree
	}
	return tree.Left.getSuccesor()
}

// // Delete from tree
// func (tree *Tree) Delete(target int) {
// 	if tree == nil {
// 		return
// 	}

// 	// search for the value's node
// 	child := tree.depthFirstSearch(target)

// 	//CASE 1: LEAF
// 	if child.isLeaf() {
// 		tree = nil
// 		return
// 	}
// 	//CASE 2: HAVE ONE CHILD
// 	if (tree.Right == nil && tree.Left != nil) || (tree.Left == nil && tree.Right != nil) {
// 		if tree.Left != nil {
// 			parent := tree.getParent(target)
// 			parent.Left = tree.Left
// 		}
// 		if tree.Right != nil {
// 			parent := tree.getParent(target)
// 			parent.Right = tree.Right
// 		}
// 		return
// 	}
// 	//CASE 3: HAVE TWO CHILDERN
// 	succesor := tree.getSuccesor()
// 	parent := tree.getParent(target)
// 	parent.Left = succesor

// }

// PRE-ORDER
func (tree *Tree) walkPreOrder(path *[]any) {
	if tree == nil {
		return
	}
	*path = append(*path, tree.Val)
	tree.Left.walkPreOrder(path)
	tree.Right.walkPreOrder(path)
}

func (tree *Tree) PreOrder() []any {
	path := []any{}
	tree.walkPreOrder(&path)
	return path
}

// IN-ORDER
func (tree *Tree) walkInOrder(path *[]any) {
	if tree == nil {
		return
	}
	tree.Left.walkInOrder(path)
	*path = append(*path, tree.Val)
	tree.Right.walkInOrder(path)
}

func (tree *Tree) InOrder() []any {
	path := []any{}
	tree.walkInOrder(&path)
	return path
}

// POST-ORDER
func (tree *Tree) walkPostOrder(path *[]any) {
	if tree == nil {
		return
	}
	tree.Left.walkPostOrder(path)
	tree.Right.walkPostOrder(path)
	*path = append(*path, tree.Val)
}

func (tree *Tree) PostOrder() []any {
	path := []any{}
	tree.walkPostOrder(&path)
	return path
}

// func (tree *Tree) Insert(value int) bool {
// 	if value == tree.Val {
// 		return false
// 	}
// 	if value > tree.Val {
// 		if tree.Right == nil {
// 			tree.Right = &Tree{Val: value}
// 			return true
// 		}
// 		return tree.Right.Insert(value)
// 	}
// 	if tree.Left == nil {
// 		tree.Left = &Tree{Val: value}
// 		return true
// 	}
// 	return tree.Left.Insert(value)
// }

// // Delete removes value from the tree and returns the resulting subtree root.
// // Callers must reassign, e.g. tree = tree.Delete(value), since the root itself
// // may change (or become nil when the last node is removed).
// func (tree *Tree) Delete(value int) *Tree {
// 	if tree == nil {
// 		return nil
// 	}
// 	if value < tree.Val {
// 		tree.Left = tree.Left.Delete(value)
// 		return tree
// 	}
// 	if value > tree.Val {
// 		tree.Right = tree.Right.Delete(value)
// 		return tree
// 	}

// 	// Found the node to delete.
// 	if tree.Left == nil {
// 		return tree.Right
// 	}
// 	if tree.Right == nil {
// 		return tree.Left
// 	}

// 	// Two children: replace with the in-order successor (smallest value in the
// 	// right subtree), then delete that successor from the right subtree.
// 	successor := tree.Right
// 	for successor.Left != nil {
// 		successor = successor.Left
// 	}
// 	tree.Val = successor.Val
// 	tree.Right = tree.Right.Delete(successor.Val)
// 	return tree
// }
