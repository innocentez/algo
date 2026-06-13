package var_1

import "fmt"

var examples = [][]int{
	{1, 2, 3, 1},
	{1, 2, 3, 4},
	{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
}

func Problem() {
	for _, example := range examples {
		res := containsDuplicate(example)
		fmt.Println(res)
	}
}

func containsDuplicate(nums []int) bool {
	dupMap := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := dupMap[num]; ok {
			return true
		}
		dupMap[num] = struct{}{}
	}

	return false
}
