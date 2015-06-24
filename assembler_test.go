package main

import "testing"

func TestTokenize(t *testing.T) {
	in := "JMP LABEL ;comment here"
	want := []string{"JMP", "LABEL", ";comment", "here"}
	got := Tokenize(in)
	if got != want {
		t.Errorf("Tokenize(%q) == %q, want %q", in, got, c.want)
	}
}
