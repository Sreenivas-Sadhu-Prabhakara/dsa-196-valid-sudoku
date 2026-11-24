package main

import "fmt"

// IsValidSudoku reports whether the FILLED cells of a 9×9 board (digits
// '1'-'9', '.' for empty) violate no Sudoku rule. It does not check
// solvability — only the placed digits. Bitmasks per row, column and box.
func IsValidSudoku(board [9][9]byte) bool {
	var rows, cols, boxes [9]uint16
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			ch := board[r][c]
			if ch == '.' {
				continue
			}
			if ch < '1' || ch > '9' {
				return false
			}
			bit := uint16(1) << (ch - '1')
			b := (r/3)*3 + c/3
			if rows[r]&bit != 0 || cols[c]&bit != 0 || boxes[b]&bit != 0 {
				return false
			}
			rows[r] |= bit
			cols[c] |= bit
			boxes[b] |= bit
		}
	}
	return true
}

func main() {
	board := [9][9]byte{}
	for r := range board {
		for c := range board {
			board[r][c] = '.'
		}
	}
	board[0][0], board[0][4] = '5', '5'
	fmt.Println(IsValidSudoku(board)) // two 5s in row 0
}
