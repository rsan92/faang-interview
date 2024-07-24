package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// LeetCode: https://leetcode.com/problems/two-sum/

func TestTwoSum(t *testing.T) {
	t.Parallel()

	t.Run("Given input=[1,3,7,9,2] and target=11, Function should return [3, 4]", func(t *testing.T) {
		input := []int{1, 3, 7, 9, 2}
		target := 11

		result := TwoSum(input, target)
		expectedResult := []int{3, 4}

		assert.Equal(t, expectedResult, result)

	})

	t.Run("Given input=[] and target=11, Function should return nil", func(t *testing.T) {
		input := []int{}
		target := 11

		result := TwoSum(input, target)

		assert.Empty(t, result)

	})

	t.Run("Given input=[1] and target=11, Function should return nil", func(t *testing.T) {
		input := []int{1}
		target := 11

		result := TwoSum(input, target)

		assert.Empty(t, result)

	})

	t.Run("Given input=[1, 2, 3] and target=11, Function should return nil", func(t *testing.T) {
		input := []int{1, 2, 3}
		target := 11

		result := TwoSum(input, target)

		assert.Empty(t, result)

	})
}

func BenchmarkTwoSum(b *testing.B) {
	input := make([]int, 0)
	limit := 10000
	for i := 0; i <= limit; i++ {
		input = append(input, i)
	}
	target := limit + limit - 1

	for i := 0; i < b.N; i++ {
		TwoSum(input, target)
	}
}
