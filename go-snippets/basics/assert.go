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

func AssertErrorEquals(actualError, expectedError error, t testing.TB) {
	t.Helper()
	if actualError == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if actualError != expectedError {
		t.Errorf("expected %q but got %q", expectedError, actualError)
	}
}

func AssertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("did not want an error but got one: %q", err)
	}
}
