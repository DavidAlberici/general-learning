package dependency_injection

import (
	"bytes"
	asserts "hello-world"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	actual := buffer.String()
	expected := "Hello, Chris"

	asserts.AssertStringEquals(actual, expected, t)
}
