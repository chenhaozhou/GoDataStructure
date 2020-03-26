package BaseArray

import (
	"../Merger"
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	merger := Merger.NewAdd()
	segTree := NewSegmentTree([]int{1, 2, 3, 4, 5}, merger)
	formatE := []int{15, 3, 12, 1, 2, 3, 9, 0, 0, 0, 0, 0, 0, 4, 5, 0, 0, 0, 0, 0}
	for k, v := range formatE {
		if v != segTree.element[k] {
			t.Error("NewSegmentTree Error")
		}
	}
}

func TestSegmentTree_Query(t *testing.T) {
	merger := Merger.NewAdd()
	segTree := NewSegmentTree([]int{1, 2, 3, 4, 5, 6}, merger)
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
		t.Error("query error:2 3")
	}
	if 6 != segTree.Query(5, 5) {
		t.Error("query error:5 5")
	}
	if 5 != segTree.Query(4, 4) {
		t.Error("query error:4 4")
	}
	if 4 != segTree.Query(3, 3) {
		t.Error("query error:3 3")
	}
	if 2 != segTree.Query(1, 1) {
		t.Error("query error:1 1")
	}
	if 1 != segTree.Query(0, 0) {
		t.Error("query error:0 0 ")
	}
	if 11 != segTree.Query(4, 5) {
		t.Error("query error:4 5")
	}
}

func TestSegmentTree_Set(t *testing.T) {
	merger := Merger.NewAdd()
	segTree := NewSegmentTree([]int{1, 2, 3, 4, 5, 6}, merger)
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
