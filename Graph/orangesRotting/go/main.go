package main

import (
	"github.com/OYE0303/DSA/goutils/queue"
)

func OrangesRotting(grid [][]int) int {
	type point struct {
		row int
		col int
	}

	q := []point{}
	freshTotal, minute := 0, 0
	offsets := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for r := range grid {
		for c := range grid[r] {
			// add rotten oranges to queue
			if grid[r][c] == 2 {
				q = append(q, point{r, c})
			}

			// count fresh oranges
			if grid[r][c] == 1 {
				freshTotal++
			}
		}
	}

	// BFS
	// continue until queue is empty or no fresh oranges left
	for len(q) > 0 && freshTotal > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			p := q[0]
			q = q[1:]

			// check all 4 directions
			row, col := p.row, p.col
			for _, o := range offsets {
				r, c := row+o.row, col+o.col

				// check if the orange is out of grid or not fresh
				if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != 1 {
					continue
				}

				grid[r][c] = 2
				q = append(q, point{r, c})
				freshTotal--
			}
		}

		minute++
	}

	if freshTotal != 0 {
		return -1
	}

	return minute
}

func main() {
	// test
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	println(OrangesRotting(grid)) // 4

}

func OrangesRotting1(grid [][]int) int {
	type point struct {
		row int
		col int
	}

	q := queue.Constructor[point]()
	minute, freshTotal := 0, 0
	offsets := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for r := range grid {
		for c := range grid[r] {
			// add rotten oranges to queue
			if grid[r][c] == 2 {
				q.Push(point{r, c})
			}

			// count fresh oranges
			if grid[r][c] == 1 {
				freshTotal++
			}
		}
	}

	// BFS
	// continue until queue is empty or no fresh oranges left
	for q.Size() > 0 && freshTotal > 0 {
		size := q.Size()
		for i := 0; i < size; i++ {
			p := q.Front()
			q.Pop()

			// check all 4 directions
			row, col := p.row, p.col
			for _, o := range offsets {
				r, c := row+o.row, col+o.col

				// check if the orange is out of grid or not fresh
				if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != 1 {
					continue
				}

				grid[r][c] = 2
				q.Push(point{r, c})
				freshTotal--
			}
		}

		minute++
	}

	if freshTotal != 0 {
		return -1
	}

	return minute
}
