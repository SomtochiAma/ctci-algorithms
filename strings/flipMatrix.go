package strings

import (
	"fmt"
)


// Time complexity O(N^2)
// Space complexity O(N^2)
func reverseMatrix(matrix [][]int) [][]int{
	newMatrix := [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			col := i
			row := Abs(len(matrix) - 1 - j)
			fmt.Println(i, j, col, row)
			newMatrix[row][col] = matrix[i][j]
		}
	}
	
	return newMatrix
}


// Time complexity O(N^2)
// Space complexity O(1)
func flipPicture(m [][]int) [][]int {
	n := len(m)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 -layer
		
		for i := first; i < last; i++ {
			offset := i - first
			
			top := m[first][i]
			
			m[first][i] = m[last-offset][first]
			m[last-offset][first] = m[last][last-offset]
			m[last][last-offset] = m[i][last]
			m[i][last] = top
		}
	}
	
	return m

}

func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}
