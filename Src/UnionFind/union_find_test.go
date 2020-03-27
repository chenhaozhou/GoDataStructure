package UnionFind

import (
	"math/rand"
	"testing"
)

func TestNewUnion(t *testing.T) {
	uf := NewUnionFind(10)
	if uf.IsConnected(0, 3) {
		t.Error("0,3 not connected")
	}

	uf.Connect(0, 3)
	if !uf.IsConnected(0, 3) {
		t.Error("0,3 is connected")
	}

	if uf.IsConnected(0, 4) {
		t.Error("0,4 not connected")
	}

	uf.Connect(0, 4)
	if !uf.IsConnected(3, 3) {
		t.Error("3,3 is connected")
	}
	if !uf.IsConnected(0, 4) {
		t.Error("0,4 is connected")
	}

	if uf.IsConnected(6, 7) {
		t.Error("6,7 not connected")
	}
	uf.Connect(6, 7)

	if !uf.IsConnected(6, 7) {
		t.Error("6,7 is connected")
	}
}

func TestUnionFind_Connect(t *testing.T) {
	size := 10000
	uf := NewUnionFind(size)
	for i := 0; i < size; i++ {
		uf.Connect(i, i)
	}

	for i := 0; i < size; i++ {
		uf.IsConnected(rand.Intn(size), rand.Intn(size))
	}
}

func BenchmarkNewUnion(b *testing.B) {
	uf := NewUnionFind(b.N)
	for i := 0; i < b.N; i++ {
		uf.Connect(rand.Intn(b.N), rand.Intn(b.N))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uf.IsConnected(rand.Intn(b.N), rand.Intn(b.N))
	}
}
