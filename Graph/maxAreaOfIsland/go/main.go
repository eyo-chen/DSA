package main

import (
	"fmt"

	"github.com/OYE0303/DSA/goutils/queue"
)

// use DFS(Recursion, Stack)
func MaxAreaOfIsland(grid [][]int) int {
	seen := map[string]bool{}
	maxCount := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				continue
			}

			key := fmt.Sprintf("%d-%d", r, c)
			if _, ok := seen[key]; ok {
				continue
			}

			count := helper(grid, r, c, seen)
			if count > maxCount {
				maxCount = count
			}
		}
	}

	return maxCount
}

func helper(grid [][]int, r int, c int, seen map[string]bool) int {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
		return 0
	}

	if grid[r][c] == 0 {
		return 0
	}

	key := fmt.Sprintf("%d-%d", r, c)
	if _, ok := seen[key]; ok {
		return 0
	}

	seen[key] = true

	return 1 + helper(grid, r+1, c, seen) + helper(grid, r-1, c, seen) + helper(grid, r, c-1, seen) + helper(grid, r, c+1, seen)
}

// use BFS(Queue)
func MaxAreaOfIsland1(grid [][]int) int {
	seen := map[string]bool{}
	maxCount := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				continue
			}

			key := fmt.Sprintf("%d-%d", r, c)
			if _, ok := seen[key]; ok {
				continue
			}

			count := helperQueue(grid, r, c, seen)
			if count > maxCount {
				maxCount = count
			}
		}
	}

	return maxCount
}

type pair struct {
	r int
	c int
}

func helperQueue(grid [][]int, r int, c int, seen map[string]bool) int {
	q := queue.Constructor[pair]()
	q.Push(pair{r, c})

	count := 0
	for !q.Empty() {
		p := q.Front()
		q.Pop()

		if p.r < 0 || p.c < 0 || p.r >= len(grid) || p.c >= len(grid[0]) {
			continue
		}

		if grid[r][c] == 0 {
			continue
		}

		key := fmt.Sprintf("%d-%d", p.r, p.c)
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = true
		count++

		q.Push(pair{r + 1, c})
		q.Push(pair{r - 1, c})
		q.Push(pair{r, c + 1})
		q.Push(pair{r, c - 1})
	}

	return count
}
