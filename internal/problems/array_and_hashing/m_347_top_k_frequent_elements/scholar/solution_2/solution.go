package solution_2

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

func Solution() {
	//res := topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)
	//res := topKFrequent([]int{2, 2, 2, 3, 3, 2}, 2)
	//res := topKFrequent([]int{3, 0, 1, 0}, 1)
	//res := topKFrequent([]int{1, 1, 1, 2, 2, 3333}, 2)
	res := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(res)
}

func topKFrequent(nums []int, k int) []int {
	slices.Sort(nums)
	result := make([]int, 0, k)
	compMap := make(map[int][]int)

	used := 0
	for i := 0; i < len(nums); i++ {
		used++

		if i+1 == len(nums) || nums[i] != nums[i+1] {
			compMap[used] = append(compMap[used], nums[i])
			used = 0
		}
	}

	req := k

	keys := slices.Collect(maps.Keys(compMap))
	slices.SortFunc(keys, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	for _, key := range keys {
		if len(result) == k {
			break
		}

		if req > len(compMap[key]) {
			result = append(result, compMap[key]...)
			req -= len(compMap[key])
		} else {
			result = append(result, compMap[key][:req]...)
		}
	}

	return result
}
