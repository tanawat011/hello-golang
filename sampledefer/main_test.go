package main

import (
	"bytes"
	"testing"
)

func TestDeferredCallRunsLast(t *testing.T) {
	var buf bytes.Buffer
	run(&buf)
	if got, want := buf.String(), "first\nsecond\nfinish\n"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
