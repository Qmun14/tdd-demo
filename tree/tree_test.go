package tree_test

import (
	"testing"

	. "github.com/Qmun14/tdd-demo/tree"
	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	nd := NewNode(0, nil, nil)
	assert.NotNil(t, nd)
}

func TestIsBinarySearchTree_Nil(t *testing.T) {
	var node *Node
	assert.True(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_OneNode(t *testing.T) {
	node := NewNode(5, nil, nil)
	assert.True(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_TwoLevel_Fail(t *testing.T) {
	node := NewNode(1, NewNode(2, nil, nil), NewNode(3, nil, nil))
	assert.False(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_TwoLevel_LeftNotExist(t *testing.T) {
	node := NewNode(1, nil, NewNode(3, nil, nil))
	assert.True(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_TwoLevel_RightNotExist(t *testing.T) {
	node := NewNode(2, NewNode(1, nil, nil), nil)
	assert.True(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_TwoLevel_Valid(t *testing.T) {
	node := NewNode(2, NewNode(1, nil, nil), NewNode(3, nil, nil))
	assert.True(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_Invalid_ifDuplicateLeft(t *testing.T) {
	node := NewNode(2, NewNode(2, nil, nil), NewNode(2, nil, nil))
	assert.False(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_Invalid_ifDuplicateRight(t *testing.T) {
	node := NewNode(2, NewNode(1, nil, nil), NewNode(2, nil, nil))
	assert.False(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_InValid_ThreeLevel(t *testing.T) {
	node := NewNode(6, NewNode(3, NewNode(5, nil, nil), NewNode(4, nil, nil)), NewNode(8, nil, nil))
	assert.False(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_valid_ThreeLevel(t *testing.T) {
	node := NewNode(6, NewNode(3, NewNode(2, nil, nil), NewNode(5, nil, nil)), NewNode(8, nil, nil))
	assert.True(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_Valid_ThreeLevel(t *testing.T) {
	node := NewNode(6,
		NewNode(3, nil, nil),
		NewNode(8, NewNode(7, nil, nil), NewNode(9, nil, nil)))
	assert.True(t, IsBinarySearchTree(node))
}
