package sudoku

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

const (
	rowSize    = 9
	columnSize = 9
)

var numbList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

var gridSolution [][]int

func CreateGrid() [][]int {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}

	fillGrid(&grid)
	gridSolution = grid
	fmt.Println(gridSolution)

	return getGridToSolve(grid)
}

func getGridToSolve(grid [][]int) [][]int {
	//A higher number of attempts will end up removing more numbers from the grid
	//Potentially resulting in more difficult grids to solve!
	attempts := 5
	counter := 1

	for attempts > 0 {
		// Select a random cell that is not already empty
		row := rand.Intn(rowSize)
		col := rand.Intn(columnSize)

		for grid[row][col] == 0 {
			row = rand.Intn(rowSize)
			col = rand.Intn(columnSize)
		}

		//Remember its cell value in case we need to put it back
		backup := grid[row][col]
		grid[row][col] = 0

		//Take a full copy of the grid
		gridCopy := grid

		//Count the number of solutions that this grid has (using a backtracking approach implemented in the solveGrid() function)
		counter = 0
		solveGrid(&gridCopy, &counter)

		//If the number of solution is different from 1 then we need to cancel the change by putting the value we took away back in the grid
		if counter != 1 {
			grid[row][col] = backup

			//We could stop here, but we can also have another attempt with a different cell just to try to remove more numbers
			attempts = -1
		}
	}

	return grid
}

func checkGridFull(grid [][]int) bool {
	for row := 0; row < rowSize; row++ {
		for column := 0; column < rowSize; column++ {
			if grid[row][column] == 0 {
				return false
			}
		}
	}

	return true
}

func solveGrid(grid *[][]int, counter *int) bool {
	// find next empty cell
	var row, column int
	for i := 0; i < 81; i++ {
		row = i / 9
		column = i % 9

		if (*grid)[row][column] == 0 {
			for value := 0; value < len(numbList); value++ {
				// check if this value is not already used in this row
				if !has((*grid)[row], numbList[value]) {
					// check if this value is not already used in this column
					if !checkCol((*grid), column, numbList[value]) {
						// identify which of the 9 squares we are looking for
						var square [][]int
						if row < 3 {
							if column < 3 {
								square = getSquare((*grid), 0, 3, 0, 3)
							} else if column < 6 {
								square = getSquare((*grid), 0, 3, 3, 6)
							} else {
								square = getSquare((*grid), 0, 3, 6, 9)
							}
						} else if row < 6 {
							if column < 3 {
								square = getSquare((*grid), 3, 6, 0, 3)
							} else if column < 6 {
								square = getSquare((*grid), 3, 6, 3, 6)
							} else {
								square = getSquare((*grid), 3, 6, 6, 9)
							}
						} else {
							if column < 3 {
								square = getSquare((*grid), 6, 9, 0, 3)
							} else if column < 6 {
								square = getSquare((*grid), 6, 9, 3, 6)
							} else {
								square = getSquare((*grid), 6, 9, 6, 9)
							}
						}

						// Check that this value has not been already used on this 3x3 square
						if !squareHas(square, numbList[value]) {
							(*grid)[row][column] = numbList[value]

							if checkGridFull((*grid)) {
								(*counter)++
								break
							} else {
								if solveGrid(grid, counter) {
									return true
								}
							}
						}
					}
				}
			}

			break
		}
	}

	(*grid)[row][column] = 0
	return false
}

func fillGrid(grid *[][]int) bool {
	// find next empty cell
	var row, column int
	for i := 0; i < 81; i++ {
		row = i / 9
		column = i % 9

		if (*grid)[row][column] == 0 {
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(numbList), func(i, j int) { numbList[i], numbList[j] = numbList[j], numbList[i] })

			for value := 0; value < len(numbList); value++ {
				// check if this value is not already used in this row
				if !has((*grid)[row], numbList[value]) {
					// check if this value is not already used in this column
					if !checkCol((*grid), column, numbList[value]) {
						// identify which of the 9 squares we are looking for
						var square [][]int
						if row < 3 {
							if column < 3 {
								square = getSquare((*grid), 0, 3, 0, 3)
							} else if column < 6 {
								square = getSquare((*grid), 0, 3, 3, 6)
							} else {
								square = getSquare((*grid), 0, 3, 6, 9)
							}
						} else if row < 6 {
							if column < 3 {
								square = getSquare((*grid), 3, 6, 0, 3)
							} else if column < 6 {
								square = getSquare((*grid), 3, 6, 3, 6)
							} else {
								square = getSquare((*grid), 3, 6, 6, 9)
							}
						} else {
							if column < 3 {
								square = getSquare((*grid), 6, 9, 0, 3)
							} else if column < 6 {
								square = getSquare((*grid), 6, 9, 3, 6)
							} else {
								square = getSquare((*grid), 6, 9, 6, 9)
							}
						}

						// Check that this value has not been already used on this 3x3 square
						if !squareHas(square, numbList[value]) {
							(*grid)[row][column] = numbList[value]

							if checkGridFull(*grid) {
								return true
							} else {
								if fillGrid(grid) {
									return true
								}
							}
						}
					}
				}
			}

			break
		}
	}

	(*grid)[row][column] = 0
	return false
}

func has(array []int, value int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}

	return false
}

func checkCol(grid [][]int, col, value int) bool {
	for i := 0; i < rowSize; i++ {
		if grid[i][col] == value {
			return true
		}
	}

	return false
}

func checkRow(grid [][]int, row, value int) bool {
	for i := 0; i < rowSize; i++ {
		if grid[row][i] == value {
			return true
		}
	}

	return false
}

func getSquare(grid [][]int, minRow, maxRow, minCol, maxCol int) [][]int {
	var square [][]int
	square = make([][]int, 3)
	for i := minRow; i < maxRow; i++ {
		square[i%3] = grid[i][minCol:maxCol]
	}

	return square
}

func squareHas(square [][]int, value int) bool {
	for i := 0; i < len(square); i++ {
		if has(square[i], value) {
			return true
		}
	}

	return false
}

func CheckValue(grid [][]int, row, col, value int) bool {
	minRow := row / 3 * 3
	maxRow := (row/3 + 1) * 3
	minCol := col / 3 * 3
	maxCol := (col/3 + 1) * 3
	return !checkCol(grid, col, value) && !checkRow(grid, row, value) && !squareHas(getSquare(grid, minRow, maxRow, minCol, maxCol), value)
}

// GetImpossibleNum returns de values that are not possible in a given position
func GetImpossibleNum(grid [][]int, row, col int) []int {
	var availableOptions []int
	for i := 1; i < 10; i++ {
		if !CheckValue(grid, row, col, i) {
			availableOptions = append(availableOptions, i)
		}
	}

	return availableOptions
}

func CheckSolution(grid [][]int) bool {
	return reflect.DeepEqual(grid, gridSolution)
}
