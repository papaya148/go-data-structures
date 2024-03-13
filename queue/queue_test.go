package queue

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(5)
	q.Enqueue(10)
	log.Println(q.ToString())
	q.Dequeue()
	log.Println(q.ToString())
	q.Enqueue(100)
	log.Println(q.ToString())
}
