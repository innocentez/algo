package sol_2

import (
	"fmt"
	"unicode"
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

// TODO
func isPalindrome(s string) bool {
	var reverse string
	var reference string

	for i := len(s) - 1; i >= 0; i-- {
		symbol := rune(s[i])
		if unicode.IsLetter(symbol) || unicode.IsNumber(symbol) {
			reverse += string(symbol)
			reference += string(s[len(s)-i-1])
		}
		continue
	}

	fmt.Println("Reverse:", reverse, "Reference:", reference)
	return reverse == reference
}
