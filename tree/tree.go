package tree

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{val: val}
}

func (t *TreeNode) Insert(elem int) error {
	if t.val == elem {
		return errors.New("данный элемент уже имеется")
	}
	if t.val > elem {
		if t.left == nil {
			t.left = &TreeNode{val: elem}
			return nil
		}
		return t.left.Insert(elem)
	}

	if t.val < elem {
		if t.right == nil {
			t.right = &TreeNode{val: elem}
			return nil
		}
		return t.right.Insert(elem)
	}

	return nil
}

func (t *TreeNode) FindMin() int {
	if t.left == nil {
		return t.val
	}

	return t.left.FindMin()
}

func (t *TreeNode) FindMax() int {
	if t.right == nil {
		return t.val
	}

	return t.right.FindMax()
}

func (t *TreeNode) PrintInOrder() {
	if t == nil {
		return
	}

	t.left.PrintInOrder()
	fmt.Println(t.val)
	t.right.PrintInOrder()
}
