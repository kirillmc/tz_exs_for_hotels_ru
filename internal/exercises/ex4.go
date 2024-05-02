package exercises

import (
	"fmt"
	"strconv"
)

func PrintMultiplicationTable(num int) {
	resArr := setUpMatrix(num)

	for i := 1; i < num+1; i++ {
		for j := 1; j < num+1; j++ {
			resArr[i][j] = strconv.Itoa(i * j)
		}
	}

	printMatrix(resArr)
}

func setUpMatrix(num int) [][]string {
	resArr := make([][]string, num+1)
	for i := 0; i < num+1; i++ {
		resArr[i] = make([]string, num+1)
	}

	resArr[0][0] = " "

	for i := 1; i < num+1; i++ {
		resArr[i][0] = strconv.Itoa(i)
		resArr[0][i] = strconv.Itoa(i)
	}

	return resArr
}

func printMatrix(matrix [][]string) {
	for i, elem := range matrix {
		for j := range elem {
			fmt.Printf("%s\t", matrix[i][j])
		}
		fmt.Println()
	}
}
