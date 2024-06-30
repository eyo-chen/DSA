package main

import "github.com/OYE0303/DSA/goutils/queue"

type Location struct {
	row int
	col int
}

// put all the locations into the queue, then check validation
func NearestExit(maze [][]byte, entrance []int) int {
	var steps int
	q := queue.Constructor[Location]()
	q.Push(Location{row: entrance[0], col: entrance[1]})

	for !q.Empty() {
		qSize := q.Size()

		for i := 0; i < qSize; i++ {
			location := q.Front()
			q.Pop()

			r, c := location.row, location.col

			// check if the location is valid
			if r < 0 || c < 0 ||
				r >= len(maze) ||
				c >= len(maze[0]) ||
				maze[r][c] != '.' {
				continue
			}

			// check if the location is the exit
			if (r != entrance[0] || c != entrance[1]) &&
				(r == 0 || c == 0 ||
					r == len(maze)-1 ||
					c == len(maze[0])-1) {
				return steps
			}

			q.Push(Location{row: r + 1, col: c})
			q.Push(Location{row: r - 1, col: c})
			q.Push(Location{row: r, col: c + 1})
			q.Push(Location{row: r, col: c - 1})

			maze[r][c] = '-'
		}

		steps++
	}

	return -1
}

// check validation first, then put all the validate locations into the queue
func NearestExit1(maze [][]byte, entrance []int) int {
	steps := 1
	q := queue.Constructor[Location]()
	q.Push(Location{row: entrance[0], col: entrance[1]})

	offsets := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for !q.Empty() {
		qSize := q.Size()

		for i := 0; i < qSize; i++ {
			location := q.Front()
			q.Pop()

			row, col := location.row, location.col

			for _, o := range offsets {
				r, c := row+o[0], col+o[1]

				// check if the location is valid
				if r < 0 || c < 0 ||
					r >= len(maze) ||
					c >= len(maze[0]) ||
					maze[r][c] != '.' {
					continue
				}

				// check if the location is the exit
				if (r != entrance[0] || c != entrance[1]) &&
					(r == 0 || c == 0 ||
						r == len(maze)-1 ||
						c == len(maze[0])-1) {
					return steps
				}

				q.Push(Location{row: r, col: c})
				maze[r][c] = '-'
			}
		}

		steps++
	}

	return -1
}

// follow the pattern as the second solution, but not use the queue
// instead, use the slice to store the locations
func NearestExit2(maze [][]byte, entrance []int) int {
	steps := 1
	q := []Location{{entrance[0], entrance[1]}}

	offsets := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(q) > 0 {
		qSize := len(q)

		for i := 0; i < qSize; i++ {
			location := q[0]
			row, col := location.row, location.col

			// pop the first element from the queue
			// q[1:] creates a new slice that starts from the second element (index 1) of the original slice q and includes all subsequent elements.
			// q = q[1:] assigns this new slice back to q, effectively removing the first element.
			q = q[1:]

			for _, o := range offsets {
				r, c := row+o[0], col+o[1]

				// check if the location is valid
				if r < 0 || c < 0 ||
					r >= len(maze) ||
					c >= len(maze[0]) ||
					maze[r][c] != '.' {
					continue
				}

				// check if the location is the exit
				if (r != entrance[0] || c != entrance[1]) &&
					(r == 0 || c == 0 ||
						r == len(maze)-1 ||
						c == len(maze[0])-1) {
					return steps
				}

				// push the location into the queue
				q = append(q, Location{r, c})
				maze[r][c] = '-'
			}
		}

		steps++
	}

	return -1
}
