package BaseArray

import "../Merger"

type merger Merger.Merger

type segmentTree struct {
	size    int
	element []int
	merger  merger
}

func NewSegmentTree(element []int, merger merger) *segmentTree {
	n := len(element)
	t := &segmentTree{size: len(element), element: make([]int, n*4), merger: merger}
	createSegmentTree(t, 0, element)
	return t
}

func (t *segmentTree) GetSize() int {
	return t.size
}

func (t *segmentTree) IsEmpty() bool {
	return 0 == t.size
}

func (t *segmentTree) Query(L int, R int) int {
	return query(t, 0, 0, t.size-1, L, R)
}

func (t *segmentTree) Set(idx int, e int) {
	t.set(0, 0, t.size-1, idx, e)
}

func (t *segmentTree) set(rootIdx int, OL int, OR int, setIdx int, e int) {
	if OL > setIdx || OR < setIdx {
		panic("Parameter range is out of bounds!")
	}
	if OR == OL {
		t.element[rootIdx] = e
		return
	}

	mid := OL + (OR-OL-1)/2
	leftIdx := leftChildIdx(rootIdx)
	rightIdx := rightChildIdx(rootIdx)
	if setIdx > mid {
		t.set(rightIdx, mid+1, OR, setIdx, e)
	} else if setIdx <= mid {
		t.set(leftIdx, OL, mid, setIdx, e)
	}

	t.element[rootIdx] = t.merger.Merge(t.element[leftIdx], t.element[rightIdx])
}

func query(t *segmentTree, p int, OL int, OR int, L int, R int) int {
	if R < L {
		panic("Parameter range is out of bounds!")
	}
	mid := OL + (OR-OL-1)/2
	if OL == L && OR == R {
		return t.element[p]
	} else if L > mid {
		return query(t, rightChildIdx(p), mid+1, OR, L, R)
	} else if R <= mid {
		return query(t, leftChildIdx(p), OL, mid, L, R)
	} else {
		LVal := query(t, leftChildIdx(p), OL, mid, L, mid)
		RVal := query(t, rightChildIdx(p), mid+1, OR, mid+1, R)
		return t.merger.Merge(LVal, RVal)
	}

}

func createSegmentTree(t *segmentTree, p int, c []int) int {
	cLen := len(c)

	if cLen < 1 {
		return 0
	}

	if cLen == 1 {
		t.element[p] = c[0]
		return t.element[p]
	}
	t.element[p] = t.merger.Merge(createSegmentTree(t, leftChildIdx(p), c[:cLen/2]), createSegmentTree(t, rightChildIdx(p), c[cLen/2:]))
	return t.element[p]
}

func leftChildIdx(idx int) int {
	return 2*idx + 1
}

func rightChildIdx(idx int) int {
	return 2*idx + 2
}
