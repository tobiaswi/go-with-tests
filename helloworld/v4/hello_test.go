package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tobi", "")
		want := "Hello, Tobi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in their language", func(t *testing.T) {
		got := Hello("Tobi", "German")
		want := "Hallo, Tobi"
		assertCorrectMessage(t, got, want)
	})

}
