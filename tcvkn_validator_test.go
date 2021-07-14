package tcvknvalidator

import (
	"fmt"
	"testing"
)

const invalidTCVKN = 9850209056
const validTCVKN = 8160506659

func ExampleValidateTCVKN() {
	fmt.Println(ValidateTCVKN(validTCVKN))
	// Output:
	// true
}

func TestValidateTCVKN(t *testing.T) {
	t.Run("TCVKN: ok", func(t *testing.T) {
		result := ValidateTCVKN(validTCVKN)
		if result == false {
			t.Fatalf(`ValidateTCVKN(%d) %t`, validTCVKN, result)
		}
	})
	t.Run("TCVKN: failed", func(t *testing.T) {
		result := ValidateTCVKN(invalidTCVKN)
		if result == false {
			t.Fatalf(`ValidateTCVKN(%d) %t`, invalidTCVKN, result)
		}
	})

}
