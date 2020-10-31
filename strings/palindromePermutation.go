package strings

import (
	"strings"
)

/*
Notes:
- Ask if string is in ASCII or Unicode
- Other consideration: remove every non-letter
*/

// Uses an additional data structure(a hashmap)
// Space complexity is O(n) as map grows linearly with number of inputs
// Time complexity is O(n) 

func palindromePermutation(str string) bool {
	str = strings.ToLower(str)
	m := createStringMap(str)
	return checkOnlyOneOdd(m)
}

func createStringMap(str string) map[rune]int {
	m := make(map[rune]int)

	for _, char := range str {
		if string(char) == " " {
			continue
		}
		if _,ok := m[char]; ok {
			m[char]++
		} else {
			m[char] = 1
		}
	}

	return m
}


func checkOnlyOneOdd(m map[rune]int) bool{
	last := false
	
	for _, freq := range m {
		if freq % 2 != 0 {
			if last {
				return false
			}
			last = true
		}
	}
	
	return true
}