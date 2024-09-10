package utils

import "fmt"

type Queue struct{ elements []int }

func NewQueue() *Queue { return &Queue{elements: make([]int, 0)} }

func (q *Queue) Enqueue(e int) { q.elements = append(q.elements, e) }

func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		return -1
	}
	e := q.elements[0]
	q.elements = q.elements[1:]
	return e
}
func (q *Queue) Size() int     { return len(q.elements) }
func (q *Queue) IsEmpty() bool { return len(q.elements) == 0 }

func (q *Queue) Peek() int {
	if len(q.elements) == 0 {
		return -1
	}
	return q.elements[0]
}

func (q *Queue) Clear() {
	q.elements = make([]int, 0)
}

func DemoQueue() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("Queue:", *q)
	fmt.Println("Dequeue:", q.Dequeue())
	fmt.Println("Queue:", *q)
	fmt.Println("Peek:", q.Peek())
	fmt.Println("Queue:", *q)
	fmt.Println("Size:", q.Size())
	fmt.Println("IsEmpty:", q.IsEmpty())
	fmt.Println("Clear")
	q.Clear()
	fmt.Println("Queue:", *q)
	fmt.Println("Size:", q.Size())
	fmt.Println("IsEmpty:", q.IsEmpty())

}
