package sorting



func mergeSort(arr1, arr2 []string) []string{
	ret := []string{}
	i := 0
	j := 0
	for i < len(arr1) || j < len(arr2) {
		if i >= len(arr1) {
			ret = append(ret, arr2[j:]...)
			break
		}
		if j >= len(arr2) {
			ret = append(ret, arr1[i:]...)
			break
		}
		if arr1[i] < arr2[j] {
			ret = append(ret, string(arr1[i]))
			i++
		} else {
			ret = append(ret, string(arr2[j]))
			j++
		}
				
	}

	return ret
}