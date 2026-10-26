package main

import (
	"fmt"
	"trees"
)

func main() {
	tree := trees.New(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(3)

	fmt.Println("===============PRE-ORDER===============")
	fmt.Println(tree.PreOrder())
	fmt.Println("===============IN-ORDER===============")
	fmt.Println(tree.InOrder())
	fmt.Println("===============POST-ORDER===============")
	fmt.Println(tree.PostOrder())
}
