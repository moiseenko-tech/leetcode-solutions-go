package p0001_two_sum

import "sort"

func twoSum(nums []int, target int) []int {
	N := len(nums)

	// positions in the original slice
	indexes := make([]int, N)
	for i := 0; i < N; i++ {
		indexes[i] = i
	}

	// sort nums indexes
	sort.Slice(indexes, func(i, j int) bool {
		return nums[indexes[i]] < nums[indexes[j]]
	})

	// perform a search in sorted slice
	i, j := 0, N-1
search:
	for i < j {
		sum := nums[indexes[i]] + nums[indexes[j]]

		switch {
		case sum > target:
			j--
		case sum < target:
			i++
		default:
			break search
		}
	}

	return []int{indexes[i], indexes[j]}
}
