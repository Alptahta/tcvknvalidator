package tcvknvalidator

func ValidateTCVKN(tcvkn int) bool {

	if NumberOfDecimalDigits(tcvkn) != 10 {
		return NumberOfDecimalDigits(tcvkn) == 10
	}

	return true
}

func NumberOfDecimalDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func IntegerToSlice(tcvkn int) []int {
	a := make([]int, 10)
	return a
}
