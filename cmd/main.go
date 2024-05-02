package main

import (
	"fmt"

	"tz_exs_for_hotels_ru/internal/exercises"
)

func main() {

	fmt.Printf("EX1 for 3 = %v\n\n", exercises.PrintComputersCount(3))

	fmt.Printf("EX2 for [42, 12, 18] = %v\n\n", exercises.GetArrayOfCommonDivisors([]int{42, 12, 18}))

	fmt.Printf("EX3 for [11,20] = %v\n\n", exercises.GetPrimeNumbers(11, 20))

	fmt.Print("EX4 for 5 = \n")
	exercises.PrintMultiplicationTable(5)
}
