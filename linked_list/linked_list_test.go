package linkedlist_test

import (
	"testing"

	. "github.com/Qmun14/tdd-demo/linked_list"
	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	nd := NewNode(0)
	assert.NotNil(t, nd)
}

func TestNewSet(t *testing.T) {
	t.Run("A LinkedList is empty on construction", func(t *testing.T) {
		myList := NewLinkedList()
		assert.True(t, myList.IsEmpty())
	})

	t.Run("A LinkedList has size 0 on construction", func(t *testing.T) {
		myList := NewLinkedList()
		assert.Equal(t, 0, myList.Length)
	})
}

func TestInsertMyList(t *testing.T) {

	t.Run("After n push to an empty LinkedList (n > 0), the LinkedList is non-empty and its size equals n", func(t *testing.T) {
		myList := NewLinkedList()
		node1 := NewNode(12)
		node2 := NewNode(9)
		node3 := NewNode(32)
		myList.Prepend(node1)
		myList.Prepend(node2)
		myList.Prepend(node3)

		assert.NotEmpty(t, myList)
		assert.Equal(t, 3, myList.Length)
	})

	t.Run("If one pushes x then peeks, the value returned is x, and the size stay the same", func(t *testing.T) {
		myList := NewLinkedList()
		node1 := NewNode(10)
		node2 := NewNode(13)
		node3 := NewNode(2)
		myList.Prepend(node1)
		myList.Prepend(node2)
		myList.Prepend(node3)

		assert.NotEmpty(t, myList)
		peek1, _ := myList.Peeks()
		assert.Equal(t, 2, peek1)
		assert.Equal(t, 3, myList.Length)
	})

	t.Run("If one pops non-Empty LinkedList and the value popped is x, and the size is one less than old size", func(t *testing.T) {
		myList := NewLinkedList()
		node1 := NewNode(45)
		node2 := NewNode(27)
		node3 := NewNode(16)
		node4 := NewNode(32)
		node5 := NewNode(11)
		node6 := NewNode(8)
		myList.Prepend(node1)
		myList.Prepend(node2)
		myList.Prepend(node3)
		myList.Prepend(node4)
		myList.Prepend(node5)
		myList.Prepend(node6)

		assert.Equal(t, 6, myList.Length)
		myList.Pop(8)
		assert.Equal(t, 5, myList.Length)
		myList.Pop(16)
		assert.Equal(t, 4, myList.Length)

	})

	t.Run("If one pops the value that doesnt exist in List, the size stay the same, and return an ErrNoSuchElement", func(t *testing.T) {
		myList := NewLinkedList()
		node1 := NewNode(45)
		node2 := NewNode(27)
		node3 := NewNode(16)
		node4 := NewNode(32)
		myList.Prepend(node1)
		myList.Prepend(node2)
		myList.Prepend(node3)
		myList.Prepend(node4)

		err := myList.Pop(100)
		assert.Error(t, err)
		assert.Equal(t, 4, myList.Length)
	})

	t.Run("Peeking into an empty List  return an exception: ErrNoSuchElement", func(t *testing.T) {
		emptyList := NewLinkedList()
		_, err := emptyList.Peeks()
		assert.Error(t, err)
	})

	t.Run("Popping from an empty List return an error: ErrNoSuchElement", func(t *testing.T) {
		emptyList := NewLinkedList()
		err := emptyList.Pop(2)
		assert.Error(t, err)
	})
}
