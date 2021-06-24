package tcvknvalidator_test

import (
	"strconv"
	"testing"
)

func lenLoop(i int) int {
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

func lenItoa(i int) int {
	return len(strconv.Itoa(i))
}

const num = 834589

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lenLoop(num)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lenItoa(num)
	}
}
