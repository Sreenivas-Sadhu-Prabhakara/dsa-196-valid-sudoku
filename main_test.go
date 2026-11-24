package main

import "testing"

func fromRows(rows [9]string) [9][9]byte {
	var b [9][9]byte
	for r, row := range rows {
		copy(b[r][:], row)
	}
	return b
}

func TestIsValidSudoku(t *testing.T) {
	valid := fromRows([9]string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	})
	if !IsValidSudoku(valid) {
		t.Error("classic valid board rejected")
	}

	// Same board with an 8 duplicated in the first column's box.
	invalid := valid
	invalid[0][0] = '8' // clashes with 8 at row 3 col 0 (same column)
	if IsValidSudoku(invalid) {
		t.Error("column duplicate accepted")
	}

	// Box duplicate without row/column duplicate.
	boxDup := fromRows([9]string{
		"12.......",
		"..1......",
		".........",
		".........", ".........", ".........", ".........", ".........", ".........",
	})
	if IsValidSudoku(boxDup) {
		t.Error("box duplicate accepted")
	}

	empty := fromRows([9]string{
		".........", ".........", ".........", ".........", ".........",
		".........", ".........", ".........", ".........",
	})
	if !IsValidSudoku(empty) {
		t.Error("empty board is trivially valid")
	}
}
