package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 wod2 word3 word4\n")
	expected := 4
	result := count(b, false)

	if result != expected {
		t.Errorf("Expected: %d, got %d instead\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 wod2 word3\nline2\nline3 word1")
	expected := 3
	result := count(b, true)

	if result != expected {
		t.Errorf("Expected: %d, got %d instead\n", expected, result)
	}
}
