package AvlTree

import (
	"math"
)

type node struct {
	e      int
	height int
	left   *node
	right  *node
}

type avlTree struct {
	size int
	root *node
}

func newNode(e int) *node {
	return &node{e: e, height: 1}
}

func NewAvlTree() *avlTree {
	return &avlTree{size: 0}
}

func getHeight(n *node) int {
	if nil == n {
		return 0
	}
	return n.height
}

func getBalanceFactor(n *node) int {
	return getHeight(n.left) - getHeight(n.right)
}

func (t *avlTree) GetSize() int {
	return t.size
}

func (t *avlTree) IsEmpty() bool {
	return 0 == t.GetSize()
}

func (t *avlTree) Add(e int) {
	var ok bool
	t.root, ok = Add(t.root, e)
	if ok {
		t.size++
	}
}

//宏观语义：给 n 为根节点的 AvlTree 添加 e ，并返回添加后的根节点
func Add(n *node, e int) (*node, bool) {
	var ok bool
	if nil == n {
		return newNode(e), true
	}
	if n.e < e {
		n.right, ok = Add(n.right, e)
	} else if n.e > e {
		n.left, ok = Add(n.left, e)
	} else {
		n.e = e
		ok = false
	}
	//维护avl的平衡性
	n = keepBalance(n)
	return n, ok
}

func (t *avlTree) Delete(e int) {
	var ok bool
	t.root, ok = del(t.root, e)
	if ok {
		t.size--
	}
}

func del(n *node, e int) (*node, bool) {
	var rootNode *node
	ok := false
	if nil != n {
		if e == n.e {
			ok = true
			if nil == n.left && nil == n.right {
				rootNode = nil
			} else if nil == n.left {
				rootNode = n.right
			} else if nil == n.right {
				rootNode = n.left
			} else {
				root, minNode := delMin(n.right)
				minNode.right = root
				minNode.left = n.left
				rootNode = minNode
			}
		} else {
			if e < n.e {
				n.left, ok = del(n.left, e)
			} else {
				n.right, ok = del(n.right, e)
			}
			rootNode = keepBalance(n)
		}
	}
	return rootNode, ok
}

func (t *avlTree) DeleteMin() *node {
	var res *node
	t.root, res = delMin(t.root)
	if nil != res {
		t.size--
	}
	return res
}

func delMin(n *node) (*node, *node) {
	var rootNode, minNode *node
	if nil != n {
		//该节点还不是最小节点
		if nil != n.left {
			n.left, minNode = delMin(n.left)
			rootNode = keepBalance(n)
		} else {
			//该节点已经是最小节点，并且存在右子树
			if nil != n.right {
				rootNode, minNode = n.right, n
				//该节点已经是最小节点，不存在右子树
			} else {
				rootNode, minNode = nil, n
			}
		}
	}
	return rootNode, minNode
}

func keepBalance(n *node) *node {
	n.height = int(math.Max(float64(getHeight(n.left)), float64(getHeight(n.right)))) + 1
	balanceFactor := getBalanceFactor(n)
	//平衡因子绝对值大于1，则需要通过旋转维持平衡
	if balanceFactor > 1 && getBalanceFactor(n.left) >= 0 {
		return rightRotate(n)
	} else if balanceFactor < -1 && getBalanceFactor(n.right) <= 0 {
		return leftRotate(n)
	} else if balanceFactor > 1 && getBalanceFactor(n.left) < 0 {
		n.left = leftRotate(n.left)
		return rightRotate(n.right)
	} else if balanceFactor < -1 && getBalanceFactor(n) > 0 {
		n.right = rightRotate(n.right)
		return leftRotate(n)
	}
	return n
}

//右旋转
//  		y                          x
//	       / \                       /   \+

//        x   T4					/     \
//       / \         ------>	   z       y
//      z   T3					  / \     / \
//     / \						 T1  T2  T3  T4
//    T1  T2
func rightRotate(y *node) *node {
	x := y.left
	t := x.right

	x.left = y
	y.right = t

	y.height = int(math.Max(float64(getHeight(y.right)), float64(getHeight(y.left)))) + 1
	x.height = int(math.Max(float64(getHeight(x.right)), float64(getHeight(x.left)))) + 1

	return x
}

//左旋转										 x
//				y                              /   \
//			   / \							  /     \
//			  T4  x							 y       Z
//				 / \		------>			/ \     / \
//		   		T3  z					   T4  T3  T2 T1
//				   / \
//				  T2  T1
func leftRotate(y *node) *node {
	x := y.right
	t := x.left

	x.left = y
	y.right = t

	y.height = int(math.Max(float64(getHeight(y.left)), float64(getHeight(y.right)))) + 1
	x.height = int(math.Max(float64(getHeight(x.left)), float64(getHeight(x.right)))) + 1

	return x
}
