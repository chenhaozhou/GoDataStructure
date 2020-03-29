package AvlTree

import (
	"math"
	"testing"
)

func TestNewAvlTree(t *testing.T) {
	avl := NewAvlTree()
	for i := 0; i < 12; i++ {
		avl.Add(i)
	}
	if 12 != avl.GetSize() {
		t.Error("get size error")
	}

	if !isBst(avl) {
		t.Error("is not bst")
	}

	if !isAvl(avl.root) {
		t.Error("is not avl")
	}

}

func TestAvlTree_DeleteMin(t *testing.T) {
	avl := NewAvlTree()
	for i := 0; i < 12; i++ {
		avl.Add(i)
	}
	if 0 != avl.DeleteMin().e {
		t.Error("delete min error")
	}
	if 11 != avl.GetSize() {
		t.Error("deleted size not change")
	}
	if !isBst(avl) {
		t.Error("is not bst")
	}
	if !isAvl(avl.root) {
		t.Error("is not avl")
	}
}

func TestAvlTree_Delete(t *testing.T) {
	avl := NewAvlTree()
	for i := 0; i < 12; i++ {
		avl.Add(i)
	}
	if !isBst(avl) {
		t.Error("is not bst")
	}
	if !isAvl(avl.root) {
		t.Error("is not avl")
	}
	avl.DeleteMin()
	if !isBst(avl) {
		t.Error("is not bst")
	}
	if !isAvl(avl.root) {
		t.Error("is not avl")
	}
	avl.DeleteMin()
	if !isBst(avl) {
		t.Error("is not bst")
	}
	if !isAvl(avl.root) {
		t.Error("is not avl")
	}
	avl.DeleteMin()
	if !isBst(avl) {
		t.Error("is not bst")
	}
	if !isAvl(avl.root) {
		t.Error("is not avl")
	}
	avl.DeleteMin()
	if !isBst(avl) {
		t.Error("is not bst")
	}
	if !isAvl(avl.root) {
		t.Error("is not avl")
	}
	avl.Delete(5)

	if !isBst(avl) {
		t.Error("is not bst")
	}
	if !isAvl(avl.root) {
		t.Error("is not avl")
	}
	avl.Delete(7)

	if !isBst(avl) {
		t.Error("is not bst")
	}
	if !isAvl(avl.root) {
		t.Error("is not avl")
	}
}

func isBst(avl *avlTree) bool {
	var nums []int
	inOrder(avl.root, &nums)
	for i := 1; i < avl.GetSize(); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func inOrder(n *node, nums *[]int) {
	if n == nil {
		return
	}
	inOrder(n.left, nums)
	*nums = append(*nums, n.e)
	inOrder(n.right, nums)
}

func isAvl(n *node) bool {
	if nil == n {
		return true
	}
	balance := getBalanceFactor(n)
	if math.Abs(float64(balance)) > 1 {
		return false
	}
	return isAvl(n.left) && isAvl(n.right)
}
