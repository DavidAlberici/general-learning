package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeats character 5 times using strings library", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"

		assertOperation(actual, expected, t)
	})

	t.Run("Repeats character 6 times using for loop", func(t *testing.T) {
		actual := RepeatWithFor("a", 6)
		expected := "aaaaaa"

		assertOperation(actual, expected, t)
	})
}

func assertOperation(actual, expected string, t testing.TB) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	b.Run("strings library", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 15)
		}
	})

	b.Run("for loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RepeatWithFor("a", 15)
		}
	})

	b.Run("for loop that uses strings.Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RepeatWithForUsingJoin("a", 15)
		}
	})
}
