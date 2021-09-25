package arrays

type Queue struct {
	Q []interface{}
}

func (q *Queue) Push(item interface{}) {
	q.Q = append(q.Q, item)
}

func (q *Queue) Pop() interface{} {
	x := q.Q[0]
	q.Q = q.Q[1:]
	return x
}
