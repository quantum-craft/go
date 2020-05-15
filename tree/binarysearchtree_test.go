package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	var root *node = nil

	insert(&root, 8)
	insert(&root, 10)
	insert(&root, 4)
	insert(&root, 2)
	insert(&root, 6)
	insert(&root, 20)
	insert(&root, 3)
	insert(&root, 5)
	insert(&root, 2)
	insert(&root, -10)
	insert(&root, -20)

	fmt.Println(search(root, 6))
	fmt.Println(search(root, 20))
	fmt.Println(search(root, -10))
	fmt.Println(search(root, 8))
	fmt.Println(search(root, 3))
	fmt.Println(search(root, 30))
	fmt.Println(search(root, 17))
	fmt.Println(search(root, 0))
	fmt.Println(search(root, 1))

	insert(&root, 30)
	insert(&root, 17)
	insert(&root, 0)
	insert(&root, 1)

	fmt.Println(search(root, 30))
	fmt.Println(search(root, 17))
	fmt.Println(search(root, 0))
	fmt.Println(search(root, 1))
}
