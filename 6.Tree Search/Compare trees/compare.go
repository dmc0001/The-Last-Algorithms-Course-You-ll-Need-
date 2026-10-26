package main

import (
	"fmt"
	"trees"
)

func compare(a *trees.Tree, b *trees.Tree) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return compare(a.Left, b.Left) && compare(a.Right, b.Right)

}
func main() {
	// both nil
	fmt.Println(compare(nil, nil)) // true

	// one nil
	fmt.Println(compare(trees.New(1), nil)) // false

	// identical trees
	a := &trees.Tree{Val: 1, Left: &trees.Tree{Val: 2}, Right: &trees.Tree{Val: 3}}
	b := &trees.Tree{Val: 1, Left: &trees.Tree{Val: 2}, Right: &trees.Tree{Val: 3}}
	fmt.Println(compare(a, b)) // true

	// different values
	c := &trees.Tree{Val: 1, Left: &trees.Tree{Val: 2}, Right: &trees.Tree{Val: 4}}
	fmt.Println(compare(a, c)) // false

	// different structure
	d := &trees.Tree{Val: 1, Left: &trees.Tree{Val: 2}}
	e := &trees.Tree{Val: 1, Right: &trees.Tree{Val: 2}}
	fmt.Println(compare(d, e)) // false
}
