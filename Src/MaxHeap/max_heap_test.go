package MaxHeap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()
	for i := 0; i < 10; i++ {
		h.Add(rand.Intn(100))
	}

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = h.ExtractMax()
	}

	for i := 1; i < 10; i++ {
		if arr[i-1] < arr[i] {
			fmt.Println(arr[i-1], arr[i], i)
			t.Error("error")
		}
	}
}
