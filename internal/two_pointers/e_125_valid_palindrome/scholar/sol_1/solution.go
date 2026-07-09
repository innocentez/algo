package sol_1

import (
	"fmt"
	"regexp"
	"strings"
)

func Solution() {
	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
	}

	for _, testCase := range testCases {
		res := isPalindrome(testCase)
		fmt.Println(res)
	}
}

func isPalindrome(s string) bool {
	var reverse string
	var reference string = s

	nonAlphaRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	reference = nonAlphaRegex.ReplaceAllString(reference, "")
	reference = strings.ToLower(reference)

	for i := len(reference) - 1; i >= 0; i-- {
		reverse += string(reference[i])
	}

	return reverse == reference
	return true
}
