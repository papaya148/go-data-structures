package stack

import (
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	st := NewStack[int]()
	st.Push(5)
	st.Pop()
	st.Push(4)
	st.Push(6)
	log.Println(st.Peek())
	st.Pop()
	log.Println(st.ToString())
}
