package tui

import "testing"

func TestGreet(t *testing.T) {
	want := "Welcome to bean-talk, an ollama TUI"
	got := Greet()
	if got != want {
		t.Errorf("want = \"%s\", got \"%s\"", want, got)
	}
}
