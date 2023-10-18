package twosum_test

import (
	"testing"

	. "github.com/Qmun14/tdd-demo/hash_tables/two_sum"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("Should return [0, 1] when given ([1,2], 3)", func(t *testing.T) {
		nums := []int{1, 2}
		target := 3
		var expected = []int{0, 1}
		assert.Equal(t, expected, TwoSum(nums, target))
	})

	t.Run("Should return [0, 2] when given ([1,2,3], 4)", func(t *testing.T) {
		nums := []int{1, 2, 3}
		target := 4
		var expected = []int{0, 2}
		assert.Equal(t, expected, TwoSum(nums, target))
	})

	t.Run("Should return [1, 2] when given ([1,2,3], 5)", func(t *testing.T) {
		nums := []int{1, 2, 3}
		target := 5
		var expected = []int{1, 2}
		assert.Equal(t, expected, TwoSum(nums, target))
	})

	t.Run("Should return nil slice -> []int{} when given ([1,2,3], 9)", func(t *testing.T) {
		nums := []int{1, 2, 3}
		target := 9
		expected := []int{}
		assert.Equal(t, expected, TwoSum(nums, target))
	})
}
