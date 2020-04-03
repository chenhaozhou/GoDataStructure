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
	var ok bool
	ok, t.root = Add(t.root, e)
	if ok {
		t.size++
	}
}

func (t *Bst) Contains(e int) bool {
	return Search(t.root, e)
}

//前序遍历
func (t *Bst) PreOrder() []int {
	var res []int
	return preOrder(t.root, res)
}

//后序遍历
func (t *Bst) LasOrder() []int {
	var res []int
	return lasOrder(t.root, res)
}

//中序遍历
func (t *Bst) MidOrder() []int {
	var res []int
	return midOrder(t.root, res)
}

//层序遍历
func (t *Bst) LevelOrder() []int {
	var q []*node
	var res []int
	q = append(q, t.root)

	for i := 0; i < t.size; i++ {
		n := q[i]
		if nil != n {
			res = append(res, q[i].e)
			if nil != n.left {
				q = append(q, n.left)
			}
			if nil != n.right {
				q = append(q, n.right)
			}
		}
	}

	return res
}

//寻找最大值
func (t *Bst) FindMax() int {
	node := t.root
	for nil != node.right {
		node = node.right
	}

	return node.e
}

//寻找最小值
func (t *Bst) FindMin() int {
	node := t.root
	for nil != node.left {
		node = node.left
	}

	return node.e
}

func (t *Bst) Delete(v int) {
	var ok bool
	ok, t.root = del(t.root, v)
	if ok {
		t.size--
	}
}

func del(n *node, v int) (bool, *node) {
	var ok bool
	if nil == n {
		return false, nil
	}

	if n.e > v {
		ok, n.left = del(n.left, v)
	} else if n.e < v {
		ok, n.right = del(n.right, v)
	} else {
		if nil == n.right {
			newLeftNode := n.left
			n = nil
			return true, newLeftNode
		} else if nil == n.left {
			newRightNode := n.right
			n = nil
			return true, newRightNode
		} else {
			//找到待删除节点右子树最小节点
			minNode, tNode := delMin(n.right)
			minNode.left = n.left
			minNode.right = tNode
			n = nil
			return true, minNode
		}
	}

	return ok, n
}

func delMin(n *node) (*node, *node) {
	var delNode *node

	if nil == n {
		return nil, nil
	}

	if nil == n.left && nil == n.right {
		return n, nil
	} else if nil == n.left {
		return n, n.right
	} else {
		delNode, n.left = delMin(n.left)
	}

	return delNode, n.left
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

func Add(n *node, e int) (bool, *node) {
	var ok bool
	if nil == n {
		return true, &node{e: e, left: nil, right: nil}
	}

	if n.e < e {
		ok, n.right = Add(n.right, e)
	} else if n.e > e {
		ok, n.left = Add(n.left, e)
	} else {
		return false, n
	}
	return ok, n
}
