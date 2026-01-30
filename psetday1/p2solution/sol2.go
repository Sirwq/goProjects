package main

func main() {

}

func diagonalSum(squareArr2D [][]int) (int, interface{}) {
	size := len(squareArr2D)
	var sum int = 0

	if size != len(squareArr2D[0]) {
		return sum, "You passed not squared array"
	}

	for i := range size {
		for j := range size {
			if i == j {
				sum += squareArr2D[i][j]
			} else if i+j == size-1 {
				sum += squareArr2D[i][j]

			}
		}
	}

	return sum, nil

}
