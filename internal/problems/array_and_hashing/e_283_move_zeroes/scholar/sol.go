package scholar

import "fmt"

func Solution() {
	testCase := []int{0, 1, 0, 3, 12}
	res := moveZeroes(testCase)
	fmt.Println(res)
}

func moveZeroes(nums []int) []int {
	origLen := len(nums) - 1

	for i, v := range nums {
		if v > 0 {
			nums = append(nums, v)
		}

		if i == origLen {
			nums = nums[origLen:]
		}
	}

	return nums
}
