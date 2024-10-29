package structs

import (
	asserts "hello-world"
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Perimeter()
		asserts.AssertFloatEquals(actual, expected, 0.000001, t)
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{6.4, 7.2}
		expected := 27.2

		checkPerimeter(t, rectangle, expected)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{6.7}
		expected := 42.0973415581

		checkPerimeter(t, circle, expected)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		asserts.AssertFloatEquals(actual, expected, 0.000001, t)
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.2, 1.1}
		expected := 5.72

		checkArea(t, rectangle, expected)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{6.7}
		expected := 141.02609422

		checkArea(t, circle, expected)
	})
}

func TestAreaWithTableTests(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 5.2, Height: 1.1}, expected: 5.72},
		{"Circle", Circle{6.7}, 141.02609422},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.shape.Area()
			asserts.AssertFloatEquals(actual, test.expected, 0.000001, t)
		})
	}
}
