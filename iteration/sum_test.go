package iteration

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expected := 15

		if actual != expected {
			t.Errorf("got %d want %d given, %v", actual, expected, numbers)
		}
	})

}
