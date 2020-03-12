package Bst

import "testing"

func create() *Bst {
	arr := [12]int{80, 32, 90, 12, 54, 32, 1, 43, 21, 5, 100, 91}
	bst := NewBst()

	for _, e := range arr {
		bst.Add(e)
	}
	return bst
}

func TestBst_Size(t *testing.T) {
	bst := NewBst()
	if bst.Size() != 0 {
		t.Error("size error")
	}
	bst.Add(1)
	bst.Add(2)
	bst.Add(3)
	bst.Add(4)
	if bst.Size() != 4 {
		t.Error("size error")
	}

	bst.Add(1)
	if bst.Size() != 4 {
		t.Error("size error")
	}

	bst.Delete(0)
	if bst.Size() != 4 {
		t.Error("size error")
	}

	bst.Delete(1)
	if bst.Size() != 3 {
		t.Error("size error")
	}

	bst.Delete(2)
	if bst.Size() != 2 {
		t.Error("size error")
	}

	bst.Delete(3)
	bst.Delete(4)
	if bst.Size() != 0 {
		t.Error("size error")
	}
}

func TestBst_Contains(t *testing.T) {
	arr := [12]int{80, 32, 90, 12, 54, 32, 1, 43, 21, 5, 100, 91}
	bst := create()
	for _, v := range arr {
		if !bst.Contains(v) {
			t.Error("bst contains error")
		}
	}
}

func TestBst_FindMin(t *testing.T) {
	bst := create()
	if 1 != bst.FindMin() {
		t.Error("bst findMin err")
	}
}

func TestBst_FindMax(t *testing.T) {
	bst := create()
	if 100 != bst.FindMax() {
		t.Errorf("bst findMax err")
	}
}

func TestBst_PreOrder(t *testing.T) {
	perOrder := [11]int{80, 32, 12, 1, 5, 21, 54, 43, 90, 100, 91}
	bst := create()
	bstPerOrder := bst.PreOrder()
	for i, v := range perOrder {
		if v != bstPerOrder[i] {
			t.Error("perOrder err")
			t.Error(bstPerOrder)
		}
	}
}

func TestBst_MidOrder(t *testing.T) {
	midOrder := [11]int{1, 5, 12, 21, 32, 43, 54, 80, 90, 91, 100}
	bst := create()
	bstMidOrder := bst.MidOrder()

	for i, v := range midOrder {
		if v != bstMidOrder[i] {
			t.Error("perOrder err")
		}
	}
}

func TestBst_LasOrder(t *testing.T) {
	lasOrder := [11]int{5, 1, 21, 12, 43, 54, 32, 91, 100, 90, 80}
	bst := create()
	bstLasOrder := bst.LasOrder()

	for i, v := range lasOrder {
		if v != bstLasOrder[i] {
			t.Error("perOrder err")
		}
	}
}

func TestBst_LevelOrder(t *testing.T) {
	levelOrder := [11]int{80, 32, 90, 12, 54, 100, 1, 21, 43, 91, 5}
	bst := create()
	bstLevelOrder := bst.LevelOrder()
	for i, v := range levelOrder {
		if v != bstLevelOrder[i] {
			t.Error("levelOrder err")
		}
	}
}

func TestBst_Delete(t *testing.T) {
	bst := create()

	bst.Delete(1000)

	lasOrder := [11]int{5, 1, 21, 12, 43, 54, 32, 91, 100, 90, 80}
	bstLasOrder := bst.LasOrder()
	for i, v := range lasOrder {
		if v != bstLasOrder[i] {
			t.Error("delete not exists val err")
		}
	}

	bst.Delete(91)

	lasOrder1 := [10]int{5, 1, 21, 12, 43, 54, 32, 100, 90, 80}
	bstLasOrder = bst.LasOrder()
	for i, v := range lasOrder1 {
		if v != bstLasOrder[i] {
			t.Error("delete val err")
		}
	}

	bst.Delete(80)

	lasOrder2 := [9]int{5, 1, 21, 12, 43, 54, 32, 100, 90}
	bstLasOrder = bst.LasOrder()
	for i, v := range lasOrder2 {
		if v != bstLasOrder[i] {
			t.Error("delete val err")
		}
	}
}
