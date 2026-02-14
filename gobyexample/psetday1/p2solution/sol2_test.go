package main

import "testing"

func TestDiagonalSum(t *testing.T) {

	arr2D := make([][]int, 5)
	for i := range arr2D {
		inSize := 5
		arr2D[i] = make([]int, inSize)

		for j := range inSize {
			arr2D[i][j] = 7
		}
	}

	arr2Dn := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	res, err := diagonalSum(arr2D)

	if res != 63 {
		t.Errorf("expected '63', got %d", res)
		t.Errorf("return message: %v", err)
	}

	res, err = diagonalSum(arr2Dn)

	if res != 25 {
		t.Errorf("expected '25', got %d", res)
		t.Errorf("return message: %v", err)
	}

}
