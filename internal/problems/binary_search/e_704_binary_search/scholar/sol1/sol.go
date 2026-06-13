package sol1

import (
	"fmt"
)

func Solution() {
	res := search([]int{1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}, 1)
	fmt.Println("Result index is", res)
}

func search(nums []int, target int) int {
	arrLength := len(nums)
	pointer := (arrLength / 2) - 1
	maxNum := arrLength - 1
	lowNum := 0

	for {
		if nums[pointer] == target {
			return pointer
		}

		if nums[pointer] < target {
			lowNum = pointer
			pointer = ((maxNum - lowNum) / 2) + pointer + 1
			continue
		}

		if nums[pointer] > target {
			maxNum = pointer
			pointer = pointer - ((maxNum - lowNum) / 2)
			continue
		}
	}

	return -1
}

// Убит нахуй
