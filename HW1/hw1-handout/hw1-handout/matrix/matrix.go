package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(lst []int, a, b int) bool {
	if lst == nil || len(lst) == 0 {
		return false
	}
	for i := 0; i < len(lst); i++ {
		if lst[i] == a {
			if i-1 >= 0 && lst[i-1] == b {
				return true
			} else if i+1 < len(lst) && lst[i+1] == b {
				return true
			}
		}
	}
	return false
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	if mat == nil || len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	row := len(mat)
	col := len(mat[0])
	res := make([][]int, col)
	for i := 0; i < col; i++ {
		res[i] = make([]int, row)
	}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			res[i][j] = mat[j][i]
		}
	}
	return res
}

// AreNeighbors returns true iff a and b are Manhattan neighbors in the 2D
// matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == a {
				if i-1 >= 0 && mat[i-1][j] == b {
					return true
				} else if i+1 < len(mat) && mat[i+1][j] == b {
					return true
				} else if j-1 >= 0 && mat[i][j-1] == b {
					return true
				} else if j+1 < len(mat[i]) && mat[i][j+1] == b {
					return true
				}
			}
		}
	}
	return false
}
