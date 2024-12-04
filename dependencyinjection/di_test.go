package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	expected := "Hello, Chris"
	actual := buffer.String()

	if actual != expected {
		t.Errorf("expected %q actual %q", expected, actual)
	}
}
