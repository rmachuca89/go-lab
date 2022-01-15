package main

import (
	"bytes"
	"testing"
)

var (
	testString = "word1 word2 word3\nline2\nline3 word1"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString(testString)

	expected := 6
	result := count(b, false, true, false)

	if result != expected {
		t.Errorf("Expected: %d, got %d instead\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString(testString)

	expected := 3
	result := count(b, true, false, false)

	if result != expected {
		t.Errorf("Expected: %d, got %d instead\n", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString(testString)

	expected := 35
	result := count(b, false, false, true)

	if result != expected {
		t.Errorf("Expected: %d, got %d instead\n", expected, result)
	}
}
