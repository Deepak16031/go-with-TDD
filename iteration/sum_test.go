package iteration

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v actual %v", expected, actual)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, actual, expected []int) {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v actual %v", expected, actual)
		}
	}

	t.Run("sum the tails of slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, actual, expected)
	})

	t.Run("sum tails of empty slices safely", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{})
		expected := []int{0, 0}
		checkSums(t, actual, expected)
	})
}
