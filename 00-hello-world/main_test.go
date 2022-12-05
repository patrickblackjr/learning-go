package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to specific person", func(t *testing.T) {
		got := Hello("Lauren", "")
		want := "Hello Lauren"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Alex", "spanish")
		want := "Hola Alex"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Alexi", "french")
		want := "Bonjour Alexi"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
