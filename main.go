package main

import (
	"fmt"
)

func mergeSort(arr1, arr2 []string) []string{
	ret := []string{}
	i := 0
	j := 0
	for i < len(arr1) || j < len(arr2) {
		fmt.Println(i, j)
		if i >= len(arr1) {
			fmt.Println("you", j)
			ret = append(ret, arr2[j:]...)
			break
		}
		if j >= len(arr2) {
			fmt.Println("me", j)
			ret = append(ret, arr1[i:]...)
			break
		}
		if arr1[i] < arr2[j] {
			fmt.Println(arr1[i])
			ret = append(ret, string(arr1[i]))
			i++
		} else {
			fmt.Println(arr2[j])
			ret = append(ret, string(arr2[j]))
			j++
		}
				
	}
	
	fmt.Println(ret)
	return ret
}



func main() {
	fmt.Println("Em" > "Best")
	mergeSort([]string{"A", "D"}, []string{"B", "C"})
}
