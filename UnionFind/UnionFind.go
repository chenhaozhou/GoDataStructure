package UnionFind

type unionFind struct {
	size     int
	elements []int
	rank     []int
}

func NewUnionFind(sizes int) *unionFind {
	UF := unionFind{size: sizes, elements: make([]int, sizes), rank: make([]int, sizes)}
	for i := 0; i < sizes; i++ {
		UF.elements[i] = i
		UF.rank[i] = 0
	}

	return &UF
}

func findRoot(uf *unionFind, e int) int {
	if e < 0 || e > uf.size-1 {
		panic("params valid")
	}

	for e != uf.elements[e] {
		e = uf.elements[e]
	}
	return e
}

//0 1 2 3 4 5 6 7 8 9
//0 1 2 3 4 5 6 7 8 9
func (uf *unionFind) Connect(p int, q int) {
	pRoot := findRoot(uf, p)
	qRoot := findRoot(uf, q)

	if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.elements[pRoot] = uf.elements[qRoot]
	} else if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.elements[qRoot] = uf.elements[pRoot]
	} else {
		uf.elements[pRoot] = uf.elements[qRoot]
		uf.rank[pRoot] = uf.rank[pRoot] + uf.rank[qRoot]
	}
}

func (uf *unionFind) IsConnected(p int, q int) bool {
	return findRoot(uf, p) == findRoot(uf, q)
}
