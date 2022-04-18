// Chapter 15 Binary Search -- Youtube Video Walkthrough
package main

import "fmt"

/*
Steps:
1) Define Node struct
2) Insert method
	- takes key that should not be already in the tree
3) Search method
	- return true if there is a node with with that value
*/

var count int

type Node struct {
	Key int
	Left *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	count++ 

	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(150)
	fmt.Println(tree) // &{100 <nil> 0xc00000c048}

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	fmt.Println(tree.Search(76)) // true
	fmt.Println(count) // 3
	fmt.Println(tree.Search(6)) // false
}