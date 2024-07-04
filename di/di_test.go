package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
  buffer := bytes.Buffer{}
  Greet(&buffer, "Cris")

  got := buffer.String()
  want := "Hello, Cris"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
