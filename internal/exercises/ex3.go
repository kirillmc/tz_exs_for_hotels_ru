package exercises

func GetPrimeNumbers(numStart, numEnd int) []int {
	resArray := []int{}

	for i := numStart; i < numEnd+1; i++ {
		isOk := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
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
