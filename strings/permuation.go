package strings

import "fmt"

func Permutation(str string) {
	permute(str, "")
}

func permute(str, prefix string) {
	if len(str) == 0 {
		fmt.Println(prefix)
	}

	for i, _ := range str {
		permute(str[:i] + str[i+1:], prefix + string(str[i]))
	}
}