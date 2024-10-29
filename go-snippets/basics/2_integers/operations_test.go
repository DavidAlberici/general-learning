package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		got := Add(1, 4)
		want := 5

		assertOperation(got, want, t)
	})
}

func assertOperation(got, want int, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d but got %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(5, 6)
	fmt.Println(sum)
	// Output: 11
}
