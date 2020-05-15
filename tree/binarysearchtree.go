package binarysearchtree

import "fmt"

type node struct {
	Value int
	Left  *node
	Right *node
}

func insert(n **node, v int) {
	if *n == nil {
		*n = &node{Value: v, Left: nil, Right: nil}
	} else if v < (*n).Value {
		insert(&(*n).Left, v)
	} else {
		insert(&(*n).Right, v)
	}
}

func print(n *node) {
	if n != nil {
		print(n.Left)
		fmt.Println(n.Value)
		print(n.Right)
	}
}

func printPreorder(n *node) {
	if n != nil {
		fmt.Println(n.Value)
		print(n.Left)
		print(n.Right)
	}
}

func printPostorder(n *node) {
	if n != nil {
		print(n.Left)
		print(n.Right)
		fmt.Println(n.Value)
	}
}

func search(n *node, v int) bool {
	if n != nil {
		if n.Value == v {
			fmt.Print("Got it: ")
			fmt.Print(n.Value)
			fmt.Print(" ")
			return true
		} else if v < n.Value {
			return search(n.Left, v)
		} else {
			return search(n.Right, v)
		}
	}

	return false
}
