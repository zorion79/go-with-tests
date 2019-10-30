package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test for any number", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("given %v, got %d, expected %d", numbers, got, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("test two arrays", func(t *testing.T) {
		got := All([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("test two arrays", func(t *testing.T) {
		got := AllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("test three arrays", func(t *testing.T) {
		got := AllTails([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
		want := []int{5, 7, 9}

		checkSums(t, got, want)
	})

	t.Run("test if one array is empty", func(t *testing.T) {
		got := AllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(t, got, want)
	})
}
