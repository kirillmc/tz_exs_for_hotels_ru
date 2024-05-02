package exercises

func GetArrayOfCommonDivisors(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	minV := getMinElem(array)

	resArray := []int{}
	for i := 2; i < minV; i++ {
		isOk := true
		for _, j := range array {
			if j%i != 0 {
				isOk = false
				break
			}
		}

		if isOk {
			resArray = append(resArray, i)
		}
	}

	return resArray
}

func getMinElem(array []int) int {
	minV := array[0]
	for _, elem := range array {
		if elem < minV {
			minV = elem
		}
	}

	return minV
}
