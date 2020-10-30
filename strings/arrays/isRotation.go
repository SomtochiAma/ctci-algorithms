package strings

import (
	"fmt"
)

func isRotation(str1, str2 string) bool{
	var point int
	
	for i,_ := range str1 {
		if str1[i] == str2[1] {
			point = i
			break
		}
	}
	
	longStr := str2 + str1[point-1:]
	fmt.Println(point, longStr)
	
	return isSubString(longStr, str1)
}

func isSubString(str1, str2 string) bool {
	return true
}

func main() {
	fmt.Println(isRotation("waterbottle", "erbottlewat"))
}
