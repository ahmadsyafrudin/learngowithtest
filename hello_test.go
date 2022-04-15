package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Rudi", "English")
		want := "Hello, Rudi"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when empty string is Supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Joan", "French")
		want := "Bonjour, Joan"
		assertCorrectMessage(t, got, want)
	})
}
