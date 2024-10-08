package asserts

import (
	"math"
	"testing"
)

func AssertStringEquals(actual, expected string, t testing.TB) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func AssertIntEquals(actual, expected int, t testing.TB) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func AssertFloatEquals(actual, expected, maxDifferenceAllowed float64, t testing.TB) {
	t.Helper()
	diff := math.Abs(actual - expected)
	if diff > maxDifferenceAllowed {
		t.Errorf("expected %f but got %f", expected, actual)
	}
}
