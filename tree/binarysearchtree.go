package binarysearchtree

type node struct {
	val   int
	left  *node
	right *node
}

func delete(n **node, val int) bool {
	ptr, leftParent, rightParent := find(n, val, nil, nil)

	if ptr == nil {
		return false
	}

	if ptr.left == nil {
		if leftParent != nil {
			leftParent.left = ptr.right
		} else if rightParent != nil {
			rightParent.right = ptr.right
		}
	} else {
		right, rightParent := findRightmost(ptr.left, nil)
		ptr.val = right.val
		if rightParent != nil {
			rightParent.right = right.left
		} else {
			ptr.left = right.left
		}
	}

	return true
}

func findRightmost(n *node, rightParent *node) (*node, *node) {
	if n.right == nil {
		return n, rightParent
	}

	return findRightmost(n.right, n)
}

func validateBST(n *node) bool {
	if n == nil {
		return true
	} else if n.left == n.right {
		return true
	}

	if n.left != nil && n.left.val > n.val {
		return false
	}

	if n.right != nil && n.right.val <= n.val {
		return false
	}

	return validateBST(n.left) && validateBST(n.right)
}

func find(n **node, val int, leftParent, rightParent *node) (*node, *node, *node) {
	if *n == nil {
		return nil, leftParent, rightParent
	}

	if (*n).val < val {
		return find(&(*n).right, val, nil, *n)
	} else if (*n).val > val {
		return find(&(*n).left, val, *n, nil)
	} else {
		return *n, leftParent, rightParent
	}
}

func insert(n **node, val int) {
	if *n == nil {
		(*n) = &node{
			val:   val,
			left:  nil,
			right: nil,
		}

		return
	}

	if (*n).val < val {
		insert(&((*n).right), val)
	} else {
		insert(&((*n).left), val)
	}
}
