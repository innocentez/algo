package scholar

import (
	"fmt"
	"maps"
	"slices"
)

func Solution() {
	testCases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
	}

	for _, testCase := range testCases {
		res := groupAnagrams(testCase)
		fmt.Println(res)
	}
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string

	var used = make([]int, 0, len(strs))
	for i := 0; i < len(strs); i++ {
		fmt.Println("prev I is", strs[i])

		if slices.Contains(used, i) {
			continue
		}

		arr := make([]string, 0)
		arr = append(arr, strs[i])
		fmt.Println("I is", strs[i])

		counts := make(map[rune]int)
		for _, r := range strs[i] {
			counts[r]++
		}

		fmt.Println("Counts is:", counts)

	OuterLoop:
		for j := i + 1; j < len(strs); j++ {
			if len(strs[i]) != len(strs[j]) {
				continue
			}

			qwerty := maps.Clone(counts)
			fmt.Println("Qwerty is:", counts)

			fmt.Println("J to compare is", strs[j])
			for _, q := range strs[j] {
				fmt.Println("Qwerty update:", qwerty)
				qwerty[q]--
				if qwerty[q] < 0 {
					continue
				}
			}

			for _, val := range qwerty {
				fmt.Println("Qwerty range:", qwerty)
				if val != 0 {
					fmt.Println("Exit from arr:", val, "Try to compare", strs[j], "with", strs[i])
					continue OuterLoop
				}
			}

			fmt.Println("Enter", strs[j], "to", arr)
			arr = append(arr, strs[j])
			used = append(used, j)
		}

		res = append(res, arr)
	}

	return res
}
