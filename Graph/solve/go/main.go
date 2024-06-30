package main

import (
	"fmt"

	"github.com/OYE0303/DSA/goutils/queue"
	"github.com/OYE0303/DSA/goutils/stack"
)

func Solve(board [][]byte) {
	safeRegion := map[string]bool{}
	lastRow := len(board) - 1
	lastCol := len(board[0]) - 1

	// check the first and last row(=)
	for r := 0; r < len(board); r++ {
		if board[r][0] == 'O' {
			helper(board, r, 0, safeRegion)
		}

		if board[r][lastCol] == 'O' {
			helper(board, r, lastCol, safeRegion)
		}
	}

	// check the first and last column(||)
	for c := 0; c < len(board[0]); c++ {
		if board[0][c] == 'O' {
			helper(board, 0, c, safeRegion)
		}

		if board[lastRow][c] == 'O' {
			helper(board, lastRow, c, safeRegion)
		}
	}

	// iterate through the board and change the 'O' to 'X' if it is not in the
	// safeRegion
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] != 'O' {
				continue
			}

			key := fmt.Sprintf("%d-%d", r, c)
			if _, ok := safeRegion[key]; ok {
				continue
			}

			board[r][c] = 'X'
		}
	}

}

func helper(board [][]byte, r int, c int, safeRegion map[string]bool) {
	// check if the location is out of bound
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return
	}

	// check if the location is already explored
	key := fmt.Sprintf("%d-%d", r, c)
	if _, ok := safeRegion[key]; ok {
		return
	}

	// check if the location is not 'O'
	// if it's not 'O', then there's no need to explore the adjacent locations
	if board[r][c] == 'X' {
		return
	}

	// mark the location as safe
	safeRegion[key] = true

	// push the adjacent locations into the data structure
	helper(board, r+1, c, safeRegion)
	helper(board, r-1, c, safeRegion)
	helper(board, r, c+1, safeRegion)
	helper(board, r, c-1, safeRegion)
}

type pair struct {
	r int
	c int
}

func helperStack(board [][]byte, r int, c int, safeRegion map[string]bool) {
	s := stack.Constructor[pair]()
	s.Push(pair{r, c})

	for !s.Empty() {
		p := s.Peak()
		s.Pop()

		if p.r < 0 || p.c < 0 || p.r >= len(board) || p.c >= len(board[0]) {
			continue
		}

		key := fmt.Sprintf("%d-%d", p.r, p.c)
		if _, ok := safeRegion[key]; ok {
			continue
		}

		if board[p.r][p.c] == 'X' {
			continue
		}

		safeRegion[key] = true

		s.Push(pair{p.r + 1, p.c})
		s.Push(pair{p.r - 1, p.c})
		s.Push(pair{p.r, p.c + 1})
		s.Push(pair{p.r, p.c - 1})
	}
}

func helperQueue(board [][]byte, r int, c int, safeRegion map[string]bool) {
	q := queue.Constructor[pair]()
	q.Push(pair{r, c})

	for !q.Empty() {
		p := q.Front()
		q.Pop()

		if p.r < 0 || p.c < 0 || p.r >= len(board) || p.c >= len(board[0]) {
			continue
		}

		key := fmt.Sprintf("%d-%d", p.r, p.c)
		if _, ok := safeRegion[key]; ok {
			continue
		}

		if board[p.r][p.c] == 'X' {
			continue
		}

		safeRegion[key] = true

		q.Push(pair{p.r + 1, p.c})
		q.Push(pair{p.r - 1, p.c})
		q.Push(pair{p.r, p.c + 1})
		q.Push(pair{p.r, p.c - 1})
	}
}
