package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	matrix [][]int
}

func New(s string) (*Matrix, error) {
	// panic("Please implement the New function")
	var matrice Matrix
	var matrixMakingSlice []string
	var errCheck error
	var totalElem int
	//list based '\n'
	result := strings.Split(s, "\n")
	for index := 0; index < len(result); index++ {
		if index == 0 {
			//1st data , so do not required to compare with previous data
			totalElem, errCheck = elemErrorChecking(&result[index])
			if errCheck != nil {
				return nil, errCheck
			}
		} else if index != 0 {
			current, errCheck := elemErrorChecking(&result[index])
			if errCheck != nil {
				return nil, errCheck
			}
			if totalElem != current {
				return nil, fmt.Errorf("rows element count is not matching")
			}
		}
		matrixMakingSlice = append(matrixMakingSlice, result[index])
	}

	matrice.matrix = matrixMaker(matrixMakingSlice)
	return &matrice, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	//
	rows := m.Rows()

	//we will perform transpose
	xl := len(rows[0])
	yl := len(rows)

	transposed := make([][]int, xl)
	for i := range transposed {
		transposed[i] = make([]int, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			transposed[i][j] = rows[j][i]
		}
	}
	return transposed
}

func (m *Matrix) Rows() [][]int {
	// panic("Please implement the Rows function")

	rows := [][]int{}
	//this row must be independent copy
	for _, value := range m.matrix {
		newRow := make([]int, len(value))
		copy(newRow, value)
		rows = append(rows, newRow)
	}
	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	matrix := m.Rows()

	xl := len(matrix[0])
	yl := len(matrix)

	if row < 0 || row > xl-1 || col < 0 || col > yl-1 {
		return false
	}
	m.matrix[row][col] = val
	return true
}

func elemErrorChecking(elements *string) (int, error) {
	//1st data , so do not required to compare with previous data
	element := strings.Split(strings.TrimSpace(*elements), " ")
	//check every element is inside or not
	for _, value := range element {
		_, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			return 0, fmt.Errorf("only int allowed")
		}

	}
	totalElem := len(element)
	return totalElem, nil
}

func matrixMaker(input []string) [][]int {
	matrix := [][]int{}

	for _, value := range input {
		var row []int
		elements := strings.Split(strings.TrimSpace(value), " ")
		for _, elem := range elements {
			intElem, _ := strconv.Atoi(elem)
			row = append(row, intElem)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
