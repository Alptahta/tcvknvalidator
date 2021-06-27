package tcvknvalidator

import "testing"

func TestTCVKNValidator(t *testing.T) {
	t.Run("TCVKN: ok", func(t *testing.T) {
		result := ValidateTCVKN(9850209056)
		if result == false {
			t.Fatalf(`ValidateTCVKN(9850209056) %t`, result)
		}
	})
	t.Run("TCVKN: failed", func(t *testing.T) {
		result := ValidateTCVKN(8160506653)
		if result == true {
			t.Fatalf(`ValidateTCVKN(8160506653) %t`, result)
		}
	})

}
