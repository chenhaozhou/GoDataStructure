package Bst

type Queue struct {
	size int
	q    []*node
}

func (q *Queue) Enqueue(v *node) {
	q.q = append(q.q, v)
	q.size++
}

func (q *Queue) Dequeue() *node {
	if q.size == 0 {
		return nil
	}
	res := q.q[0]
	for i := 0; i < q.size; i++ {
		if i == q.size-1 {
			break
		}
		q.q[i] = q.q[i+1]
	}
	q.size--
	return res
}

func (q *Queue) Size() int {
	return q.size
}
