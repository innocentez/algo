package solution_1

import "fmt"

func Solution() {
	res := topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)
	fmt.Println(res)
}

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	compMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		compMap[nums[i]]++
	}

	for key, v := range compMap {
		if k == 0 {
			break
		}

		k--
	}

	return res
}
