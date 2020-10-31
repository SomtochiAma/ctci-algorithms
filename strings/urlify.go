package strings

import (
	"strings"
)


/**
Loops over array and replaces " " => "%20"
Time complexity = O(N)
Space complexity = O(N) creating array equal to the length of the string
**/
func urlify(str string, len int) string{
	strArray := strings.Split(str, "")
	strArray = strArray[:len]
	for i, char := range strArray {
		if char == " " {
			strArray[i] = "%20"
		}
	}
	
	return strings.Join(strArray, "")
}

/**
This function assumes that it is an array of characters and you cannot set one item to more than one character
It basically does the same thing with the additional overhead of accounting for a bigger array of characters.
This isn't needed in Go but could be useful in languages like Java
Time complexity = O(N)
Space complexity = O(N) creating array equal to the length of the string
**/
func urlify2(str string, len int) string{
	strArray := strings.Split(str, "")
	count := 0
	for _, char := range strArray {
		if char == " " {
			count++
		}
	}
	newLen := len + (2*count)
	newStrArray := make([]string, newLen) 
	index := newLen - 1
	
	for i := len-1; i >= 0; i-- {
		if strArray[i] == " " {
			newStrArray[index] = "0"
			index--
			newStrArray[index] = "2"
			index--
			newStrArray[index] = "%"
			index--
		} else {
			newStrArray[index] = strArray[i]
			index--
		}
	} 
	
	return strings.Join(newStrArray, "")
}