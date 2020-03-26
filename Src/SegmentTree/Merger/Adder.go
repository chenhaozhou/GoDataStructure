package Merger

type Add struct{}

func NewAdd() *Add {
	return &Add{}
}

func (merge *Add) Merge(x int, y int) int {
	return x + y
}
