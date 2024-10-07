package main

import "testing"

func TestGetHello(t *testing.T) {
	got := GetHello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
