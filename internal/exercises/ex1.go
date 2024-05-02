package exercises

import "fmt"

func PrintComputersCount(computerCount int) string {
	if computerCount >= 0 { // не может быть отрицательное количество компьютеров
		if computerCount%10 == 1 && computerCount%100 != 11 {
			return fmt.Sprintf("%d компьютер", computerCount)
		} else if (computerCount%10 == 2 || computerCount%10 == 3 || computerCount%10 == 4) && computerCount%100 != 12 && computerCount%100 != 13 && computerCount%100 != 14 {
			return fmt.Sprintf("%d компьютера", computerCount)
		} else {
			return fmt.Sprintf("%d компьютеров", computerCount)
		}
	}

	return "\n"
}
