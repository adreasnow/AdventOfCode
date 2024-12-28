package main

func getValueOfIndex(row int, col int) int {
	curRow := 1
	curCol := 1
	code := 20151125

	for {
		curRow, curCol = nextTriangleNumberCoordinate(curRow, curCol)
		code = generateNextCode(code)
		if curRow == row && curCol == col {
			break
		}
	}

	return code
}

func generateNextCode(code int) int {
	return (code * 252533) % 33554393
}

func nextTriangleNumberCoordinate(row int, col int) (int, int) {
	if row == 1 {
		return col + 1, 1
	}
	return row - 1, col + 1
}
