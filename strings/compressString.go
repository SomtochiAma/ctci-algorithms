package strings

import (
	"strings"
	"strconv"
)

func compressString(str string)string {
	var  count int
	var compressed strings.Builder
	
	for i := 0; i<len(str); i++ {
		count++
		if i + 1 >= len(str) || str[i] != str[i+1] {
			compressed.WriteString(string(str[i]))
			compressed.WriteString(strconv.Itoa(count))
			count = 0
		}
	}
	
	if len(compressed.String()) < len(str) {
		return compressed.String()
	} 
	
	return str
}