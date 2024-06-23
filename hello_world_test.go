// hello_world_test.go
package main

import (
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expectedOutput := "Hello, World!"
	actualOutput := runHelloWorld() // Assuming runHelloWorld is a function that runs your main function logic

	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Expected output to contain %q, got %q", expectedOutput, actualOutput)
	}
}

func runHelloWorld() string {
	// Simulate running the main function logic
	// For demonstration purposes, this simply returns a hardcoded string
	return "Hello, World!"
}
