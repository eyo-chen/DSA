package main

import "strings"

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	table := make([]string, numRows)

	row := 0
	down := true
	for _, r := range s {
		table[row] += string(r)

		// if down is true, we go down one row
		// if down is false, we go up one row
		if down {
			row++
		} else {
			row--
		}

		// if row is equal to numRows, it means it's over the edge
		// we need to go up two rows and change the direction
		if row == numRows {
			row -= 2
			down = false
		}

		// if row is equal to -1, it means it's over the edge
		// we need to go down two rows and change the direction
		if row == -1 {
			row += 2
			down = true
		}
	}

	return genAns(table)
}

func Convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	table := make([]string, numRows)

	// we have to set the initial direction to up
	// because when we hit row == 0, it will change the direction to down
	down := false
	row := 0
	for _, r := range s {
		table[row] += string(r)

		// if row is equal to 0 or numRows-1, it means we hit the edge
		// we need to change the direction
		if row == 0 || row == numRows-1 {
			down = !down
		}

		// if down is true, we go down one row
		// if down is false, we go up one row
		if down {
			row++
		} else {
			row--
		}
	}

	return genAns(table)
}

func genAns(table []string) string {
	var sb strings.Builder
	for _, s := range table {
		sb.WriteString(s)
	}
	ans := sb.String()
	return ans
}

// This is the solution I came up with at 12/22/2024
func Convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	table := make([][]byte, numRows)
	ans := make([]byte, 0, len(s))
	direction := false // false = down, true = up
	row := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		table[row] = append(table[row], char)

		// When it's the first row, we need to go down
		if row == 0 {
			direction = false
			row++
			continue
		}

		// When it's the last row, we need to go up
		if row == numRows-1 {
			direction = true
			row--
			continue
		}

		// If direction is true, we go up one row
		// If direction is false, we go down one row
		if direction {
			row--
		} else {
			row++
		}
	}

	// Concatenate each string in the table to get the result
	for r := 0; r < len(table); r++ {
		for c := 0; c < len(table[r]); c++ {
			ans = append(ans, table[r][c])
		}
	}

	return string(ans)
}

// This is the deprecated approach
// It uses a 2D array to store the result to simulate the zigzag pattern
// This approach is not recommended because it's not efficient and it's hard to implement
func Convert3(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	table := make([][]byte, numRows)
	for i := range table {
		table[i] = make([]byte, len(s))
	}

	idx := 0
	r := 0
	c := 0

	for {
		for r < numRows && idx < len(s) {
			table[r][c] = s[idx]
			r++
			idx++
		}

		if idx >= len(s) {
			break
		}

		r -= 2
		c++
		for r >= 0 && idx < len(s) {
			table[r][c] = s[idx]
			r--
			c++
			idx++
		}

		r += 2
		c--

		if idx >= len(s) {
			break
		}
	}

	ans := ""
	for _, row := range table {
		for _, col := range row {
			if col != 0 {
				ans += string(col)
			}
		}
	}

	return ans
}
