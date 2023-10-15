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
