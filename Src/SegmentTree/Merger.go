package SegmentTree

type merger interface {
	merge(x int, y int) int
}

type Add struct{}

func NewAdd() *Add {
	return &Add{}
}

func (merge *Add) merge(x int, y int) int {
	return x + y
}
