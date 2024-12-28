package eazymail

import (
	"fmt"
	"sync"
)

type Message struct {
	sender    string
	recipient string
	subject   string
	body      string
}

type Stack []Message

func (s *Stack) Push(val Message) {
	*s = append(*s, val)
}

func (s *Stack) Peek() (Message, error) {
	if s.isEmpty() {

		return (*s)[len(*s)-1], fmt.Errorf("Cannot peek empty stack")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) Pop() (Message, error) {
	if s.isEmpty() {
		return Message{}, fmt.Errorf("Cannot pop empty stack")
	}
	val := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return val, nil

}
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

type ConcurrentQueue struct {
	s1 Stack
	s2 Stack
	mu sync.Mutex
}

func (q *ConcurrentQueue) enqueue(val Message) {
	q.mu.Lock()
	q.s1.Push(val)
	q.mu.Unlock()
}
func (q *ConcurrentQueue) dequeue() (Message, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.s2.isEmpty() {
		for !q.s1.isEmpty() {
			v, _ := q.s1.Pop()
			q.s2.Push(v)
		}
	}
	if q.s2.isEmpty() {
		return Message{}, fmt.Errorf("Cannot dequeue empty queue")
	}
	r, _ := q.s2.Pop()
	return r, nil
}

func NewConcurrentQueue(length, capacity int) ConcurrentQueue {
	return ConcurrentQueue{
		s1: make(Stack, length, capacity),
		s2: make(Stack, length, capacity),
	}
}
