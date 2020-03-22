package SegmentTree

import (
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	merger := NewAdd()
	segTree := NewSegmentTree([]int{1, 2, 3, 4, 5}, merger)
	formatE := []int{15, 3, 12, 1, 2, 3, 9, 0, 0, 0, 0, 0, 0, 4, 5, 0, 0, 0, 0, 0}
	for k, v := range formatE {
		if v != segTree.element[k] {
			t.Error("NewSegmentTree Error")
		}
	}
}

func TestSegmentTree_Query(t *testing.T) {
	merger := NewAdd()
	segTree := NewSegmentTree([]int{1, 2, 3, 4, 5}, merger)
	if 15 != segTree.Query(0, 4) {
		t.Error("query error")
	}
	if 3 != segTree.Query(0, 1) {
		t.Error("query error")
	}
	if 9 != segTree.Query(3, 4) {
		t.Error("query error")
	}
	if 7 != segTree.Query(2, 3) {
		t.Error("query error")
	}
}

func TestSegmentTree_Set(t *testing.T) {
	merger := NewAdd()
	segTree := NewSegmentTree([]int{1, 2, 3, 4, 5}, merger)
	segTree.Set(4, 1)
	if 11 != segTree.Query(0, 4) {
		t.Error("set error")
	}
	segTree.Set(1, 1)
	if 2 != segTree.Query(0, 1) {
		t.Error("set error")
	}
	segTree.Set(3, 3)
	if 4 != segTree.Query(3, 4) {
		t.Error("set error")
	}
}
