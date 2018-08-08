package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tobi")
	want := "Hello, Tobi"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
