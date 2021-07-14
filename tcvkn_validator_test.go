package tcvknvalidator

import "testing"

const invalidTCVKN = 9850209056
const validTCVKN = 8160506653

func TestTCVKNValidator(t *testing.T) {
	t.Run("TCVKN: ok", func(t *testing.T) {
		result := ValidateTCVKN(invalidTCVKN)
		if result == false {
			t.Fatalf(`ValidateTCVKN(9850209056) %t`, result)
		}
	})
	t.Run("TCVKN: failed", func(t *testing.T) {
		result := ValidateTCVKN(validTCVKN)
		if result == true {
			t.Fatalf(`ValidateTCVKN(8160506653) %t`, result)
		}
	})

}
