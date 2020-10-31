package strings


// Time complexity = 0(N^2)
// Space complexity = 0(N)
func zeroMatrix(m [][]int) [][]int {
	var col []int
	var row []int

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 0 {
				row = append(row, i)
				col = append(col, i)
				break
			}
		}
	}

	for i := 0; i < len(row); i++ {
		for j := 0; j < len(m[row[i]]); j++ {
			m[row[i]][j] = 0
		}
	}

	for i := 0; i < len(col); i++ {
		for j := 0; j < len(m); j++ {
			m[j][col[i]] = 0
		}
	}

	return m

}

