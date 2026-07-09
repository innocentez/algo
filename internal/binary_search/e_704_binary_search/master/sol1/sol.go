package sol1

import "fmt"

func Solution() {
	res := search([]int{2, 5}, 5)
	fmt.Println("Result index is", res)
}

func search(nums []int, target int) int {
	lowP := 0
	highP := len(nums) - 1

	if highP == 0 {
		if nums[0] != target {
			return -1
		}
		return 0
	}

	for lowP < highP {
		mid := (lowP + highP) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		} else if guess > target {
			highP = mid - 1
		} else {
			lowP = mid + 1
		}
	}
	return -1
}
