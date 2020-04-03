package MaxHeap

import "fmt"

type maxHeap struct {
	size    int
	element []int
}

func NewMaxHeap() *maxHeap {
	return &maxHeap{size: 0}
}

func (h *maxHeap) GetSize() int {
	return h.size
}

func (h *maxHeap) IsEmpty() bool {
	return 0 == h.size
}

func (h *maxHeap) Add(e int) {
	if h.size < len(h.element) {
		h.element[h.size] = e
	} else {
		h.element = append(h.element, e)
	}
	h.siftUp(h.size)
	h.size++
}

func (h *maxHeap) Show() {
	fmt.Println(h.element)
}

func (h *maxHeap) FindMax() int {
	if h.IsEmpty() {
		panic("Can not findMax when heap is empty")
	}
	return h.element[0]
}

func (h *maxHeap) ExtractMax() int {
	maxE := h.FindMax()
	h.element[0], h.element[h.GetSize()-1] = h.element[h.GetSize()-1], h.element[0]
	h.size--
	h.siftDown(0)
	return maxE
}

func parentIdx(idx int) int {
	return (idx - 1) / 2
}

func leftChildIdx(idx int) int {
	return 2*idx + 1
}

func rightChildIdx(idx int) int {
	return 2*idx + 2
}

func (h *maxHeap) siftUp(idx int) {
	for idx > 0 && h.element[parentIdx(idx)] < h.element[idx] {
		h.element[parentIdx(idx)], h.element[idx] = h.element[idx], h.element[parentIdx(idx)]
		idx = parentIdx(idx)
	}
}

func (h *maxHeap) siftDown(idx int) {
	for leftChildIdx(idx) < h.GetSize() {
		maxChildIdx := leftChildIdx(idx)
		rightChildIdx := rightChildIdx(idx)
		if rightChildIdx < h.GetSize() && h.element[rightChildIdx] > h.element[maxChildIdx] {
			maxChildIdx = rightChildIdx
		}
		if h.element[maxChildIdx] < h.element[idx] {
			return
		}

		h.element[maxChildIdx], h.element[idx] = h.element[idx], h.element[maxChildIdx]
		idx = maxChildIdx
	}
}
