package main

var (
	movement = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

func Exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, row, col int, word string, index int) bool {
	if index == len(word) {
		return true
	}

	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) ||
		board[row][col] != word[index] {
		return false
	}

	temp := board[row][col]
	board[row][col] = '#' // Mark as visited

	for _, m := range movement {
		nextRow, nextCol := row+m[0], col+m[1]
		if dfs(board, nextRow, nextCol, word, index+1) {
			return true
		}
	}

	board[row][col] = temp // Restore the original value
	return false
}
