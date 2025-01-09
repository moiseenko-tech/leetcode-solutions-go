package p0001_two_sum

import "sort"

func twoSum(nums []int, target int) []int {
	N := len(nums)

	// positions in the original slice
	indexes := make([]int, N)
	for i := range N {
		indexes[i] = i
	}

	// sort nums indexes
	sort.Slice(indexes, func(i, j int) bool {
		return nums[indexes[i]] < nums[indexes[j]]
	})

	// perform a search in sorted slice
	left, right := 0, N-1
search:
	for left < right {
		sum := nums[indexes[left]] + nums[indexes[right]]

		switch {
		case sum > target:
			right--
		case sum < target:
			left++
		default:
			break search
		}
	}

	return []int{indexes[left], indexes[right]}
}
