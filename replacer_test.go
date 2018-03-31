package replacer

import (
	"testing"
)

func TestReplacer(t *testing.T) {

	t.Run("Test Positive", func(t *testing.T) {
		expectedOne := "i-love-golang-so-much"
		expectedTwo := "i_love_golang_so_much"

		result1 := Replace("i love golang so     much", "-")
		result2 := Replace("i love \n golang     so much", "_")

		if expectedOne != result1 {
			t.Errorf("not what i expect : %s", expectedOne)
		}

		if expectedTwo != result2 {
			t.Errorf("not what i expect : %s", expectedTwo)
		}
	})
}
