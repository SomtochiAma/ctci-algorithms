package strings

import (
	// "fmt"
	"sort"
	"strings"
)

/*
Notes:
- Ask if string is in ASCII or Unicode
- 
*/

// Uses an additional data structure(a hashmap)
// Space complexity is O(n) as map grows linearly with number of inputs
// Time complexity is O(n) as we are looping over the strings once
func uniqueString(s string) bool{
	m := map[rune]bool{}
	for _, item := range s {
		if m[item] {
			return false
		}
		m[item] = true
	}
	
	return true
}

// This solution doesn't use an additional data structure
// Time Complexity is O(NlogN(from sorting the string) + N(loops)) => O(NlogN)
// Space Complexity is O(1)
// [Ignoring the additional array that is formed when sorting as that is a go thing]
func uniqueString2(s string) bool{
	strArr := strings.Split(s, "")	
	sort.Strings(strArr)
	for i:=0; i < len(strArr)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}
	return true
}