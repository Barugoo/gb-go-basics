package main

import (
	"fmt"
	"sort"
)

type Queue struct {
	queueInternal []int
}

func (q *Queue) Len() int {
	return len(q.queueInternal)
}

func (q *Queue) Less(i, j int) bool {
	return q.queueInternal[i] < q.queueInternal[j]
}

func (q *Queue) Swap(i, j int) {
	q.queueInternal[i], q.queueInternal[j] = q.queueInternal[j], q.queueInternal[i]
}

func (q *Queue) Push(elem int) {
	q.queueInternal = append(q.queueInternal, elem)
}

func (q *Queue) Pop() int {
	if len(q.queueInternal) == 0 {
		return 0
	}
	elem := q.queueInternal[0]
	q.queueInternal = q.queueInternal[1:]
	return elem
}

func main() {
	q := &Queue{}

	q.Push(4)
	q.Push(8)
	q.Push(1)
	q.Push(5)
	q.Push(0)

	sort.Sort(q)
	reversedQueue := sort.Reverse(q)

	queueLength := reversedQueue.Len()
	for i := 0; i < queueLength; i++ {
		fmt.Println(reversedQueue.(*Queue).Pop())
	}
}
