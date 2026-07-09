package scholar

import (
	"fmt"
	"strings"
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
	if len(s) != len(t) {
		return false
	}

	for _, i := range []rune(s) {
		if !strings.ContainsRune(t, i) {
			return false
		}
	}

	return true
}
