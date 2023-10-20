package set_test

import (
	"testing"

	. "github.com/Qmun14/tdd-demo/set"
	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	t.Run("Should return Length 0 when a Set on Construction", func(t *testing.T) {
		set1 := NewSet()
		assert.Equal(t, 0, set1.Length)
	})

	t.Run("Should insert a new data set and after insert the length of set is increase greather than old", func(t *testing.T) {
		set1 := NewSet()
		assert.Equal(t, 0, set1.Length)

		set1.Insert("Budi")
		assert.Equal(t, 1, set1.Length)

		set1.Insert("Karjo")
		assert.Equal(t, 2, set1.Length)
	})

	t.Run("Should search data of set, after insert if data exist return true", func(t *testing.T) {
		set1 := NewSet()
		set1.Insert("Budi")
		assert.Equal(t, 1, set1.Length)
		assert.True(t, set1.Contains("Budi"))

		set1.Insert("Karjo")
		set1.Insert("Joni")
		assert.Equal(t, 3, set1.Length)
		assert.True(t, set1.Contains("Karjo"))

	})

	t.Run("Cannot insert duplicate data set the length of set will same as old and will return error such  ErrDataExist : Error Data Already exist", func(t *testing.T) {
		set1 := NewSet()
		assert.Equal(t, 0, set1.Length)

		set1.Insert("Budi")
		assert.Equal(t, 1, set1.Length)

		err := set1.Insert("Budi")
		assert.Error(t, err)
		assert.NotEqual(t, 2, set1.Length)
	})

	t.Run("Should delete data set, after that the length of set will decrease one less than old length", func(t *testing.T) {
		set1 := NewSet()
		assert.Equal(t, 0, set1.Length)

		set1.Insert("Budi")
		set1.Insert("Karman")
		set1.Insert("Ucok")
		set1.Insert("Haris")
		set1.Insert("Dora")

		assert.Equal(t, 5, set1.Length)

		assert.True(t, set1.Contains("Ucok"))
		assert.True(t, set1.Contains("Karman"))
		assert.True(t, set1.Contains("Haris"))
		assert.True(t, set1.Contains("Dora"))

		set1.Delete("Ucok")

		assert.False(t, set1.Contains("Ucok"))
		assert.Equal(t, 4, set1.Length)
		assert.True(t, set1.Contains("Dora"))
		assert.True(t, set1.Contains("Budi"))

		set1.Delete("Budi")
		assert.False(t, set1.Contains("Budi"))
		assert.Equal(t, 3, set1.Length)

		set1.Delete("Dora")
		assert.False(t, set1.Contains("Dora"))
		assert.Equal(t, 2, set1.Length)

	})
	t.Run("Cannot Delete data which doesn't exist in set, then return error : ErrDataNotExist, such Error Data doesn't exist", func(t *testing.T) {
		set1 := NewSet()
		assert.Equal(t, 0, set1.Length)
		set1.Insert("Budi")
		set1.Insert("Karman")
		set1.Insert("Ucok")
		set1.Insert("Haris")
		set1.Insert("Dora")

		err := set1.Delete("Ucup")
		assert.Error(t, err)
		assert.Equal(t, 5, set1.Length)
	})

	t.Run("Cannot insert set data if set is full max limit", func(t *testing.T) {
		set1 := NewSet()
		assert.Equal(t, 0, set1.Length)

		set1.Insert("Person1")
		set1.Insert("Person2")
		set1.Insert("Person3")
		set1.Insert("Person4")
		set1.Insert("Person5")
		set1.Insert("Person6")
		set1.Insert("Person7")
		set1.Insert("Person8")
		set1.Insert("Person9")
		set1.Insert("Person10")

		err := set1.Insert("Person11")
		assert.Error(t, err)

		err = set1.Insert("Person12")
		assert.Error(t, err)
		set1.Insert("Person12")

		assert.NotEqual(t, 12, set1.Length)
		assert.Equal(t, 10, set1.Length)

	})
}
