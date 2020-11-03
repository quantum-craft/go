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

	fmt.Println(find(&root, 6, nil, nil))
	fmt.Println(find(&root, 20, nil, nil))
	fmt.Println(find(&root, -10, nil, nil))
	fmt.Println(find(&root, 8, nil, nil))
	fmt.Println(find(&root, 3, nil, nil))
	fmt.Println(find(&root, 30, nil, nil))
	fmt.Println(find(&root, 17, nil, nil))
	fmt.Println(find(&root, 0, nil, nil))
	fmt.Println(find(&root, 1, nil, nil))

	insert(&root, 30)
	insert(&root, 17)
	insert(&root, 0)
	insert(&root, 1)

	fmt.Println(find(&root, 30, nil, nil))
	fmt.Println(find(&root, 17, nil, nil))
	fmt.Println(find(&root, 0, nil, nil))
	fmt.Println(find(&root, 1, nil, nil))
}

func TestBinarySearchTreePrint(t *testing.T) {
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

	print(root)
	// printPostorder(root)
}
