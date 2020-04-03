package Merger

import "testing"

func TestAddMerger(t *testing.T) {
	adder := NewAdd()
	if 2 != adder.Merge(1, 1) {
		t.Error("add merger error")
	}
}
