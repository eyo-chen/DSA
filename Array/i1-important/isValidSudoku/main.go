package main

import "fmt"

// Using hash table to store the key of the row, column and block
// O(1) because the board is always 9x9
// The row and col key is straightforward, but the block key is a little bit tricky
// We basically index the block from 0 ~ 2 for both row and col
// Row has three blocks 0, 1, 2
// Col has three blocks 0, 1, 2
// Total we have 9 blocks
// For example, row = 3, col = 3, we divide row and col by 3, we get 1, 1, which is in the block (1,1)
// If row = 2, col = 1, we divide row and col by 3, we get 0, 0, which is in the block (0,0)
// Look at the diagram below to see how we decide the block
func IsValidSudoku(board [][]byte) bool {
	hashTable := map[string]bool{}
	for r := range board {
		for c := range board[r] {
			val := board[r][c]
			// Skip the empty cell
			if val == '.' {
				continue
			}

			// Create the key for the row, col and block
			rowKey := fmt.Sprintf("row-%d-%c", r, val)
			colKey := fmt.Sprintf("col-%d-%c", c, val)
			blockKey := fmt.Sprintf("block-%d-%d-%c", r/3, c/3, val)

			// If the key already exists, it means the value is duplicated
			if hashTable[rowKey] || hashTable[colKey] || hashTable[blockKey] {
				return false
			}

			// Mark the key as true
			hashTable[rowKey] = true
			hashTable[colKey] = true
			hashTable[blockKey] = true
		}
	}

	return true
}

/*
(i) represent the index of the block
row = 1, col = 2 -> (0,0)
row = 3, col = 3 -> (1,1)
row = 8, col = 6 -> (2,2)
row = 5, col = 8 -> (1,2)

       (0)      (1)       (2)
    0  1  2 | 3  4  5 | 6  7  8
(0) 1       |         |
    2       |         |
	  ----------------------------
    3       |         |
(1) 4       |         |
    5       |         |
	  ----------------------------
    6       |         |
(2) 7       |         |
    8       |         |
*/
