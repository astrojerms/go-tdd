package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("human-person")
	want := "Hello, human-person."

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
