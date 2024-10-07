package main

import "testing"

func TestGetHello(t *testing.T) {
	got := GetHello("David")
	want := "Hello, David!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
