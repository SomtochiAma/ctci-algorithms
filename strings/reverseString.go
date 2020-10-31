package strings

import "strings"

func ReverseString(str string) string{
	strArr := strings.Split(str, "")
	n := len(strArr)

	for i := 0; i < n/2; i++ {
		end := n-1-i
		strArr[i], strArr[end] = strArr[end], strArr[i]
	}

	return strings.Join(strArr, "")
}