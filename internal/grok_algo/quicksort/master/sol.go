package master

func Solution() {
	testCase := []int{0, 1, 0, 3, 12}
	moveZeroes(testCase)
}

func moveZeroes(nums []int) {
	origLen := len(nums) - 1

	for i, v := range nums {
		nums = append(nums, v)

		if i == origLen {
			nums = nums[origLen:]
		}
	}
}
