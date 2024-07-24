package arrays

func TwoSum(nums []int, target int) []int {
	result := make([]int, 0)
	hashValues := map[int]int{}

	if len(nums) < 2 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		expectedValue := target - nums[i]

		if _, ok := hashValues[expectedValue]; ok {
			result = append(result, hashValues[expectedValue])
			result = append(result, i)
			return result
		}

		hashValues[nums[i]] = i

	}

	return result
}
