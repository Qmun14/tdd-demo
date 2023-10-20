package trie_test

import (
	"testing"

	. "github.com/Qmun14/tdd-demo/trie"
	"github.com/stretchr/testify/assert"
)

func TestNewTrie(t *testing.T) {
	t.Run("Should return not nil Trie on construction and return an empty root", func(t *testing.T) {
		t1 := NewTrie()
		assert.NotNil(t, t1)
		assert.Empty(t, t1.Root)
	})
}

func TestRNWTrie(t *testing.T) {
	t.Run("Should can insert a word to a trie", func(t *testing.T) {
		t1 := NewTrie()
		err := t1.Insert("argon")
		assert.NoError(t, err)
	})
	t.Run("Should can search a word from a trie and return true if exist and false if doesn't exist", func(t *testing.T) {
		t1 := NewTrie()

		toAdd := []string{
			"aragorn",
			"aragon",
			"argon",
			"eragon",
			"eragon",
			"oregon",
			"oregano",
			"oreo",
		}
		for _, v := range toAdd {
			t1.Insert(v)
		}
		assert.True(t, t1.Search("argon"))
		assert.True(t, t1.Search("oregano"))
		assert.True(t, t1.Search("oreo"))
		assert.False(t, t1.Search("ojodumeh"))
	})

}
