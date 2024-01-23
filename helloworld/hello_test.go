package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("human-person", English)
		want := "Hello, human-person."
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello, World."
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hola, Maria' when specifying spanish.", func(t *testing.T) {
		got := Hello("Maria", Spanish)
		want := "Hola, Maria."
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Bounjour, Pierre' when specifying french.", func(t *testing.T) {
		got := Hello("Pierre", French)
		want := "Bonjour, Pierre."
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
