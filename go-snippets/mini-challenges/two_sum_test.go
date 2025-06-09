package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("twoSum", func(t *testing.T) {
		got := twoSum([]int{1, 11, 13, 14, 5, 6}, 7)
		want := [2]int{5, 0}

		assertEquals(got, want, t)
	})
}

func assertEquals(got, want [2]int, t testing.TB) {
	t.Helper()
	equals := reflect.DeepEqual(got, want)
	reversedWant := [2]int{want[1], want[0]}
	equalsReverse := reflect.DeepEqual(got, reversedWant)
	if !equals && !equalsReverse {
		t.Errorf("expected %d but got %d", got, want)
	}
}
