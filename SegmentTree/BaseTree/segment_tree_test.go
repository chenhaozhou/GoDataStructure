package BaseTree

import (
	"../Merger"
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	merger := Merger.NewAdd()
	nums := []int{1, 2, 3, 4, 5, 6}
	ST := NewSegmentTree(nums, merger)
	if 1 != ST.Query(0, 0) {
		t.Error("error:Query or NewAdd error")
	}

	if 3 != ST.Query(0, 1) {
		t.Error("error:Query or NewAdd error")
	}

	if 6 != ST.Query(0, 2) {
		t.Error("error:Query or NewAdd error")
	}

	if 5 != ST.Query(1, 2) {
		t.Error("error:Query or NewAdd error")
	}

	if 9 != ST.Query(1, 3) {
		t.Error("error:Query or NewAdd error")
	}

	if 14 != ST.Query(1, 4) {
		t.Error("error:Query or NewAdd error")
	}
	if 5 != ST.Query(4, 4) {
		t.Error("error:Query or NewAdd error")
	}
}

func TestSegTree_Set(t *testing.T) {
	merger := Merger.NewAdd()
	nums := []int{1, 2, 3, 4, 5, 6}
	ST := NewSegmentTree(nums, merger)
	ST.Set(0, 0)
	if 20 != ST.Query(0, 5) {
		t.Error("Set error:set 0 0")
	}
	ST.Set(0, 9)
	if 29 != ST.Query(0, 5) {
		t.Error("Set error: set 0 9")
	}
	if 9 != ST.Query(0, 0) {
		t.Error("Set error:set 0 0 ")
	}
	if 11 != ST.Query(0, 1) {
		t.Error("set error:set 0 0 ")
	}
}
