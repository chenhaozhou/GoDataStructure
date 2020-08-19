package main

import "fmt"

type link struct {
	count int
	head  *node
}

type node struct {
	val  int
	next *node
}

func createNotCircularLink() *link {
	l := link{
		count: 0,
		head:  nil,
	}
	var cur *node
	for i := 0; i < 10; i++ {
		n := &node{
			val:  i,
			next: nil,
		}

		if l.count == 0 {
			l.head = n
			cur = n
		} else {
			cur.next = n
			cur = n
		}
		l.count++
	}
	return &l
}

func CreateHadCircularLink() *link {
	l := &link{
		count: 0,
		head:  nil,
	}
	var cur, circularHead *node

	for i := 0; i < 10; i++ {
		n := &node{
			val:  i,
			next: nil,
		}
		if i == 11 {
			circularHead = n
		}
		if l.count == 0 {
			l.head = n
			cur = n
		} else {
			cur.next = n
			cur = cur.next
		}

		l.count++
	}
	if l.count > 0 {
		cur.next = circularHead
	}
	return l
}

func checkoutHasCircular(l *link) bool {
	if l.count == 0 {
		return false
	}

	fast, slow := l.head, l.head
	for {
		if fast == nil {
			return false
		}
		fast = runFast(fast)
		slow = runSlow(slow)

		if fast == slow {
			return true
		}
	}
}

func runFast(n *node) *node {
	if n == nil {
		return nil
	}
	if n.next == nil {
		return nil
	}

	return n.next.next
}

func runSlow(n *node) *node {
	if n == nil {
		return nil
	}
	return n.next
}

func main() {
	link := createNotCircularLink()
	fmt.Println(checkoutHasCircular(link))

	link = CreateHadCircularLink()
	fmt.Println(checkoutHasCircular(link))
}
