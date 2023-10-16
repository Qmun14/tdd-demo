package linkedlist

import "fmt"

var ErrNoSuchElement = fmt.Errorf("No such Element")

func NewNode(val int) *Node {
	return &Node{
		Data: val,
	}
}

type Node struct {
	Data int
	Next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) IsEmpty() bool {
	return l.Length == 0
}

func (l LinkedList) Peeks() (int, error) {
	if l.IsEmpty() {
		return 0, ErrNoSuchElement
	}
	return l.Head.Data, nil
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l *LinkedList) Pop(val int) error {

	if l.IsEmpty() {
		return ErrNoSuchElement
	}

	if l.Head.Data == val {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	prevToDelete := l.Head
	for prevToDelete.Next.Data != val {
		if prevToDelete.Next.Next == nil {
			return ErrNoSuchElement
		}
		prevToDelete = prevToDelete.Next
	}

	prevToDelete.Next = prevToDelete.Next.Next
	l.Length--

	return nil
}
