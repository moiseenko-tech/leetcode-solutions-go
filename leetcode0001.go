package leetcode0001

import "sort"

func twoSum(nums []int, target int) []int {
	arrLength := len(nums)

	// positions in the original slice
	indexes := make([]int, arrLength)
	for i := range arrLength {
		indexes[i] = i
	}

	// sort nums indexes
	sort.Slice(indexes, func(i, j int) bool {
		return nums[indexes[i]] < nums[indexes[j]]
	})

	// perform a search in sorted slice
	left, right := 0, arrLength-1
	for left < right {
		sum := nums[indexes[left]] + nums[indexes[right]]

		switch {
		case sum > target:
			right--
		case sum < target:
			left++
		default:
			return []int{indexes[left], indexes[right]}
		}
	}

	return []int{}
}
