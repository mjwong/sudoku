package main

import (
	"fmt"
	"testing"
)

func TestInRow(t *testing.T) {
	input := "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63.."

	mat := toGrid(input)
	printSudoku(mat)

	numList := [9][]int{{1, 2, 7, 9}, // row 1
		{1, 3, 4, 6, 7, 8, 9},
		{2, 4, 5, 6, 9},
		{2, 4, 5, 6, 7, 9},
		{1, 2, 4, 7},
		{1, 2, 3, 4, 7, 8},
		{4, 6, 7, 8, 9},
		{1, 2, 3, 5, 6, 8, 9},
		{1, 4, 7, 8, 9},
	}

	for i := 0; i < 9; i++ {
		for _, v := range numList[i] {
			if inRow(mat, i, v) {
				t.Fatalf("No. %d should not be in row %d.\n", v, i)
			}
		}
	}

	// negative test
	if !inRow(mat, 0, 3) {
		t.Fatal("No. 3 should be in row 1.\n")
	}
}

func TestInCol(t *testing.T) {
	input := "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63.."

	mat := toGrid(input)

	numList := [9][]int{{2, 4, 6, 7, 8}, // col 1
		{1, 4, 6, 7, 9},
		{1, 2, 4, 8, 9},
		{1, 3, 4, 6, 7, 9},
		{2, 3, 4, 5, 7, 8},
		{1, 2, 4, 5, 7, 9},
		{1, 5, 7, 8, 9},
		{1, 2, 4, 6, 9},
		{2, 3, 6, 7, 8, 9},
	}

	for i := 0; i < 9; i++ {
		for _, v := range numList[i] {
			if inCol(mat, i, v) {
				t.Fatalf("No. %d should not be in col %d.\n", v, i)
			}
		}
	}

	// negative test
	if !inCol(mat, 0, 9) {
		t.Fatal("No. 9 should be in col 1.")
	}
}

func TestInSqu(t *testing.T) {
	input := "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63.."

	mat := toGrid(input)

	numList := [9][]int{{1, 4, 9}, // square 1
		{1, 2, 3, 4, 6, 7, 9},
		{2, 5, 6, 7, 8, 9},
		{1, 2, 4, 6, 7, 8},
		{2, 4, 5, 7},
		{1, 2, 3, 4, 7, 9},
		{2, 4, 6, 7, 8, 9},
		{1, 3, 4, 5, 7, 8, 9},
		{1, 6, 8, 9},
	}

	lstIndex := 0
	for x := 0; x < SQ; x++ {
		for y := 0; y < SQ; y++ {
			found := false
			for i := x * SQ; i < x*SQ+SQ; i++ {
				for j := y * SQ; j < y*SQ+SQ; j++ {
					// select first empty cell to test
					if mat[i][j] == 0 {
						for _, v := range numList[lstIndex] {
							if inSqu(mat, i, j, v) {
								t.Fatalf("Cell: [%d,%d]. No. %d should not be in square %d,%d. Index = %d\n", i, j, v, x, y, lstIndex)
							}
						}
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			lstIndex++
		}
	}
}

func TestGetSqList(t *testing.T) {
	input := "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63.."

	mat := toGrid(input)

	numList := [9][]int{{1, 4, 9}, // square 1
		{1, 2, 3, 4, 6, 7, 9},
		{2, 5, 6, 7, 8, 9},
		{1, 2, 4, 6, 7, 8},
		{2, 4, 5, 7},
		{1, 2, 3, 4, 7, 9},
		{2, 4, 6, 7, 8, 9},
		{1, 3, 4, 5, 7, 8, 9},
		{1, 6, 8, 9},
	}

	lstIndex := 0
	for x := 0; x < SQ; x++ {
		for y := 0; y < SQ; y++ {
			fmt.Printf("Square %d,%d\n", x, y)
			if contains(getSqList(mat, x, y), numList[lstIndex]) {
				t.Fatalf("Possibility list %d should not be in square %d,%d.\n", numList[lstIndex], x, y)
			}
			lstIndex++
		}
	}
}

func TestGetPossibilityMat(t *testing.T) {
	input := "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63.."

	list := [][]int{
		{1, 9},
		{2, 7},
		{2, 9},
		{2, 7, 9},
		{1, 4, 9},
		{1, 3, 4, 6, 7, 9},
		{3, 4, 7},
		{1, 4, 7, 9},
		{7, 8, 9},
		{6, 9},
		{6, 7, 8, 9},
		{4},
		{4, 6, 9},
		{2, 4},
		{2, 4, 9},
		{5, 9},
		{2, 4, 6, 7},
		{4, 6, 7},
		{4, 7},
		{2, 4, 5, 7},
		{7, 9},
		{2, 7, 9},
		{1, 4, 7},
		{1, 2, 4},
		{1, 7},
		{1, 2, 4},
		{2, 4, 7, 8},
		{1, 2, 4, 8},
		{4, 7},
		{2, 4, 7},
		{1, 2, 4},
		{2, 3, 7},
		{4, 8, 9},
		{4, 7, 9},
		{4, 7, 8},
		{4, 7, 9},
		{6, 8, 9},
		{2, 6, 8},
		{6, 9},
		{2, 8, 9},
		{1, 3, 9},
		{3, 5, 8},
		{1, 5, 9},
		{1, 8, 9},
		{4, 7, 8},
		{4, 7, 9},
		{4, 7, 8},
		{1, 9},
		{8, 9},
	}

	mat := toGrid(input)

	mat2, cnt := getPossibilityMat(mat)

	if cnt != 49 {
		t.Fatalf("No. of empty cells should be %d but got %d.\n", len(list), cnt)
	}

	index := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if mat2[i][j] != nil {
				if !intArrayEquals(mat2[i][j], list[index]) {
					t.Fatalf("Arrays %v and %v should be equal but not.\n", mat2[i][j], list[index])
				} else {
					index++
				}
			}
		}
	}
}
