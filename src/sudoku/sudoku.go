package sudoku

import (
	"math/rand"
	"time"
)

const (
	rowSize    = 9
	columnSize = 9
)

var numbList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

var gridSolution [9][]int

// CreateGrid creates a new sudoku matrix, where the 0 values correspond to the empty cells
func CreateGrid(difficulty int) [][]int {
	return setDifficulty(difficulty)
}

func getGridToSolve(grid [][]int) [][]int {
	//A higher number of attempts will end up removing more numbers from the grid
	//Potentially resulting in more difficult grids to solve!
	counter := 1
	attempts := 100

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
			for _, value := range numbList {
				// check if this value is not already used in this row
				if !has((*grid)[row], value) {
					// check if this value is not already used in this column
					if !checkCol((*grid), column, value) {
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
						if !squareHas(square, value) {
							(*grid)[row][column] = value

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

			for _, value := range numbList {
				// check if this value is not already used in this row
				if !has((*grid)[row], value) {
					// check if this value is not already used in this column
					if !checkCol((*grid), column, value) {
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
						if !squareHas(square, value) {
							(*grid)[row][column] = value

							if checkGridFull(*grid) || fillGrid(grid) {
								return true
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
	for i := range array {
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
	for i := range square {
		if has(square[i], value) {
			return true
		}
	}

	return false
}

// CheckValue checks if a value is valid in a given cell
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

// CheckSolution checks if the solution to the game is the correct one
func CheckSolution(grid [][]int) bool {
	for row := 0; row < rowSize; row++ {
		for col := 0; col < columnSize; col++ {
			if gridSolution[row][col] != grid[row][col] {
				return false
			}
		}
	}

	return true
}

func setDifficulty(dif int) [][]int {
	switch dif {
	case 0:
		return difficulty(30, 50)
	case 1:
		return difficulty(30, 50)
	case 2:
		return difficulty(50, 89)
	}
	return nil
}

func difficulty(min, max int) [][]int {
	zeroes := 0
	var gridToSolve [][]int

	for zeroes < min {
		gridCopy := [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0}}

		fillGrid(&gridCopy)

		// Save the grid solution
		for i := 0; i < rowSize; i++ {
			gridSolution[i] = make([]int, len(gridCopy[i]))
			copy(gridSolution[i], gridCopy[i])
		}

		gridToSolve = getGridToSolve(gridCopy)
		zeroes = CheckZeroes(gridToSolve)

		if zeroes > max {
			zeroes = 0
		}
	}

	return gridToSolve
}

// CheckZeroes counts how many zeroes does the matrix have
func CheckZeroes(grid [][]int) int {
	counter := 0
	for row := 0; row < rowSize; row++ {
		for col := 0; col < columnSize; col++ {
			if grid[row][col] == 0 {
				counter++
			}
		}
	}

	return counter
}

// GetSolution returns the solution to the game
func GetSolution() [9][]int {
	return gridSolution
}

// ReadSudoku replaces the solution of the sudoku
func ReadSudoku(sol [9][]int) {
	// Save the grid solution
	for i := 0; i < rowSize; i++ {
		gridSolution[i] = make([]int, len(sol[i]))
		copy(gridSolution[i], sol[i])
	}
}
