package Bst

type node struct {
	e     int
	left  *node
	right *node
}

type Bst struct {
	size int
	root *node
}

func NewBst() *Bst {
	return &Bst{size: 0, root: nil}
}

func (t *Bst) IsEmpty() bool {
	return t.size == 0
}

func (t *Bst) Size() int {
	return t.size
}

func (t *Bst) Add(e int) {
	t.root = Add(t.root, e)
	t.size++
}

func (t *Bst) Search(e int) bool {
	return Search(t.root, e)
}

func (t *Bst) PreOrder() []int {
	var res []int
	return preOrder(t.root, res)
}

func (t *Bst) LasOrder() []int {
	var res []int
	return lasOrder(t.root, res)
}

func (t *Bst) MidOrder() []int {
	var res []int
	return midOrder(t.root, res)
}

func (t *Bst) FindMax() (bool, int) {
	if nil == t {
		return false, 0
	}
	node := t.root
	for nil != node.right {
		node = node.right
	}

	return true, node.e
}

func (t *Bst) FindMin() (bool, int) {
	if nil == t {
		return false, 0
	}
	node := t.root
	for nil != node.left {
		node = node.left
	}

	return true, node.e
}

func midOrder(n *node, res []int) []int {
	if nil == n {
		return res
	}

	res = midOrder(n.left, res)
	res = append(res, n.e)
	res = midOrder(n.right, res)

	return res
}

func lasOrder(n *node, res []int) []int {
	if nil == n {
		return res
	}
	res = lasOrder(n.left, res)
	res = lasOrder(n.right, res)
	res = append(res, n.e)
	return res
}

func preOrder(n *node, res []int) []int {
	if nil == n {
		return res
	}
	res = append(res, n.e)
	res = preOrder(n.left, res)
	res = preOrder(n.right, res)

	return res
}

func Search(n *node, e int) bool {
	if nil == n {
		return false
	}

	if n.e < e {
		return Search(n.right, e)
	} else if n.e > e {
		return Search(n.left, e)
	} else {
		return true
	}
}

func Add(n *node, e int) *node {
	if nil == n {
		return &node{e: e, left: nil, right: nil}
	}

	if n.e < e {
		n.right = Add(n.right, e)
	} else if n.e > e {
		n.left = Add(n.left, e)
	}
	return n
}
