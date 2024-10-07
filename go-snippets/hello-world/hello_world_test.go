package main

import "testing"

func TestGetHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := GetHello("David")
		want := "Hello, David!"

		assertMessage(got, want, t)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := GetHello("")
		want := "Hello, World!"

		assertMessage(got, want, t)
	})
}

func assertMessage(got, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
