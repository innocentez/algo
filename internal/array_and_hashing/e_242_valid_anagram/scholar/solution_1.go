package scholar

import (
	"fmt"
	"slices"
)

func Solution() {
	testCases := [][]string{
		{"eat", "tea"},
		{"anagram", "nagaram"},
		{"rat", "car"},
	}

	for _, testCase := range testCases {
		res := isAnagram(testCase[0], testCase[1])
		fmt.Println(res)
	}
}

func isAnagram(s string, t string) bool {
	isEqual := slices.Equal(sArr, tArr)
	return isEqual
}
