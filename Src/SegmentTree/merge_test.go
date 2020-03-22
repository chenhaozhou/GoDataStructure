package SegmentTree

import "testing"

func TestAddMerger(t *testing.T) {
	adder := NewAdd()
	if 2 != adder.merge(1, 1) {
		t.Error("add merger error")
	}
}
