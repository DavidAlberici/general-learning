package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	actual := Sum(numbers)
	expected := 6

	assertArraySum(actual, expected, numbers, t)
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{6, 11, 3}

	actual := SumAll(numbers1, numbers2)
	expected := []int{6, 20}

	assertArraysEqual(actual, expected, t)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum tails", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2, 7}, []int{0, 9})
		expected := []int{9, 9}

		assertArraysEqual(actual, expected, t)
	})

	t.Run("safely sum empty arrays", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}

		assertArraysEqual(actual, expected, t)
	})

	t.Run("safely sum 1 element arrays", func(t *testing.T) {
		actual := SumAllTails([]int{300}, []int{0, 65, 35})
		expected := []int{0, 100}

		assertArraysEqual(actual, expected, t)
	})
}

func assertArraySum(actual, expected int, array []int, t testing.TB) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %d but got %d given, %v", expected, actual, array)
	}
}

func assertArraysEqual(actual, expected []int, t testing.TB) {
	t.Helper()
	if !slices.Equal(actual, expected) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
