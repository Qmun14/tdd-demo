package hashtables_test

import (
	"testing"

	. "github.com/Qmun14/tdd-demo/hash_tables"
	"github.com/stretchr/testify/assert"
)

// ✅ A HashTables is not empty on construction

// ✅ A HashTables has size ArraySize on construction

// ✅ Can Insert a value which is string to an HashTables

// ✅ If one Insert x then search, the value returned true

// ✅ Cannot Insert a value which is already exist an HashTables, and  return an error ErrDataExist

// ✅ If one Delete the value is x, and after that the value should not be exist again in The hashtables

// ✅ If one Delete the value that doesnt exist in HashTables, and  return an error  ErrNoSuchElement

func TestDummy(t *testing.T) {
	hashT1 := New()
	assert.NotNil(t, hashT1)
	assert.Equal(t, 10, hashT1.Length())
}

func TestCRUD_ToHashTables(t *testing.T) {
	t.Run("Can Insert a value which is string to an HashTables", func(t *testing.T) {
		has1 := New()
		err := has1.Insert("RANDY")
		assert.NoError(t, err)
	})

	t.Run("If one insert x then search, the value returned true", func(t *testing.T) {
		has1 := New()
		has1.Insert("RANDY")

		assert.True(t, has1.Search("RANDY"))
	})

	t.Run("Cannot Insert a value which is already exist an HashTables, and return Error", func(t *testing.T) {
		has1 := New()
		err := has1.Insert("RANDY")
		assert.NoError(t, err)

		err = has1.Insert("RANDY")
		assert.Error(t, err)
	})

	t.Run("If one Delete the value is x, and after that the value should not be exist again in The Hashtables", func(t *testing.T) {
		has1 := New()
		has1.Insert("RANDY")
		has1.Insert("LUKAZ")
		has1.Insert("DAMIEN")
		has1.Insert("SHERLOCK")
		has1.Insert("JAMES")
		has1.Insert("NINA")
		has1.Insert("DRUGA")

		assert.True(t, has1.Search("SHERLOCK"))
		assert.True(t, has1.Search("NINA"))

		has1.Delete("NINA")
		assert.False(t, has1.Search("NINA"))

	})

	t.Run("If one Delete the value that doesnt exist in HashTables, and  return an ErrNoSuchElement", func(t *testing.T) {
		has1 := New()
		has1.Insert("RANDY")
		has1.Insert("LUKAZ")
		has1.Insert("SHERLOCK")
		has1.Insert("JAMES")
		has1.Insert("DRUGA")

		assert.True(t, has1.Search("RANDY"))
		assert.True(t, has1.Search("LUKAZ"))
		assert.True(t, has1.Search("SHERLOCK"))
		assert.True(t, has1.Search("JAMES"))
		assert.True(t, has1.Search("DRUGA"))

		has1.Delete("SHERLOCK")
		assert.False(t, has1.Search("SHERLOCK"))

		err := has1.Delete("ASSDDSA")
		assert.Error(t, err)

	})
}
