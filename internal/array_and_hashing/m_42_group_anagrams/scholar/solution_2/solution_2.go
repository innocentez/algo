package solution_2

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

	var compMap = make(map[string][]string, len(strs))
	for _, str := range strs {
		arr := []rune(str)
		slices.Sort(arr)
		compStr := string(arr)
		compMap[compStr] = append(compMap[compStr], str)
	}

	for v := range maps.Values(compMap) {
		res = append(res, v)
	}

	return res
}
