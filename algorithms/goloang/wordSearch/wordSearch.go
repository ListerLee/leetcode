package wordsearch

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	ick := make([][]bool, m)
	for i := 0; i < m; i++ {
		ick[i] = make([]bool, n)
	}
	var check func(x, y, k int) bool
	check = func(x, y, k int) bool {
		if board[x][y] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		ick[x][y] = true
		defer func() { ick[x][y] = false }()
		for _, dir := range directions {
			if newX, newY := x+dir.x, y+dir.y; newX >= 0 && newX < m && newY >= 0 && newY < n && ick[newX][newY] == false {
				if check(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
