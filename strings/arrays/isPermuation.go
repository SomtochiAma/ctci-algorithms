package strings

import (

	"strings"
	"sort"
)

/*
- Is the problem case sensitive?
- Is it problem space sensitive?
*/

// helper functions
func search(s string, c rune) bool{
	for _, char := range s {
		if char == c {
			return true
		}
	}
	
	return false
}


func remove(s string, c string) string{
	return strings.Replace(s, c, "", 1)
}

func sortStr(s string) string{
	strArr := strings.Split(s, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}

// Space Complexity = 0(1)
// Time Complexity = NlogN
func isPermutation(s,t string) bool{
	if len(s) != len(t) {
		return false
	}
	s = sortStr(s)
	t = sortStr(t)
	return s == t
}


// Space Complexity = 0(1)
// Time Complexity = N(loop)*N(search) N^2
func isPermutation2(s,t string) bool{
	if len(s) != len(t) {
		return false
	}
	for _, char := range s {
		exists := search(t, char)
		if !exists {
			return false
		} else {
			t = remove(t, string(char))
		}
	}
	
	return true
}
