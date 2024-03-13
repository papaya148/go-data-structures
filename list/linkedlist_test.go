package list

import (
	"log"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(5)
	list.Append(56)
	list.Append(100)
	list.Insert(6, 2)
	list.Delete(4)
	log.Println(list.ToString())
}
