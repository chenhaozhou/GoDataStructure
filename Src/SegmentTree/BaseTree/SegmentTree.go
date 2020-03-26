package BaseTree

import (
	"../Merger"
)

type merger Merger.Merger

type node struct {
	val        int
	leftChild  *node
	rightChild *node
}

type segTree struct {
	size   int
	merger merger
	root   *node
}

func NewSegmentTree(nums []int, merger merger) *segTree {
	return &segTree{size: len(nums), merger: merger, root: createSegmentTree(merger, nums)}
}

func (this *segTree) Query(l int, r int) int {
	if r < l {
		panic("params valid")
	}
	if r >= this.size {
		r = this.size - 1
	}
	return query(this.merger, this.root, 0, this.size-1, l, r)
}

func (this *segTree) Set(idx int, val int) {
	if idx < 0 || idx > this.size-1 {
		panic("params valid")
	}
	this.root = set(this.merger, this.root, 0, this.size-1, idx, val)
}

func set(merger merger, node *node, or int, ol int, idx int, val int) *node {
	if nil == node.leftChild {
		node.val = val
		return node
	}

	mid := or + (ol-or-1)/2
	if idx <= mid {
		node.leftChild = set(merger, node.leftChild, or, mid, idx, val)
	} else {
		node.rightChild = set(merger, node.rightChild, mid+1, or, idx, val)
	}

	node.val = merger.Merge(node.leftChild.val, node.rightChild.val)
	return node
}

func query(merger merger, node *node, ol int, or int, l int, r int) int {
	if ol == l && or == r {
		return node.val
	}
	mid := ol + (or-ol-1)/2
	if r <= mid {
		return query(merger, node.leftChild, ol, mid, l, r)
	} else if l > mid {
		return query(merger, node.rightChild, mid+1, or, l, r)
	} else {
		leftVal := query(merger, node.leftChild, ol, mid, l, mid)
		rightVal := query(merger, node.rightChild, mid+1, or, mid+1, r)
		return merger.Merge(leftVal, rightVal)
	}
}

func createSegmentTree(merger merger, nums []int) *node {
	if 0 == len(nums) {
		return nil
	}
	node := &node{}
	if 1 == len(nums) {
		node.val = nums[0]
	} else {
		mid := len(nums) / 2
		node.leftChild = createSegmentTree(merger, nums[:mid])
		node.rightChild = createSegmentTree(merger, nums[mid:])
		node.val = merger.Merge(node.leftChild.val, node.rightChild.val)
	}
	return node
}
