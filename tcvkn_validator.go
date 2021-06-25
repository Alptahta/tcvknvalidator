package tcvknvalidator

import (
	"fmt"
	"math"
)

func ValidateTCVKN(tcvkn int) bool {
	if NumberOfDecimalDigits(tcvkn) != 10 {
		return NumberOfDecimalDigits(tcvkn) == 10
	}
	fmt.Println(IntegerToSlice(tcvkn))
	s := IntegerToSlice(tcvkn)
	result := 0
	lastDigit := s[len(s)-1]

	for i := 0; i < len(s)-i; i++ {
		if (s[i]+9-i)%10 == 9 {
			result += 9
		} else {
			result += ((s[i] + 9 - i) % 10) * int(math.Pow(2, float64(9-i))) % 9
		}
	}

	var expectedLastDigit int = (10 - (result % 10)) % 10

	if lastDigit == expectedLastDigit {
		return true
	} else {
		return false
	}
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

func IntegerToSlice(i int) []int {
	var a []int
	for i != 0 {
		a = append(a, i%10)
		i /= 10
	}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
