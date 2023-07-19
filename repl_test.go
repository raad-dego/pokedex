package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This is a test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "    Extra    Spaces    ",
			expected: []string{"extra", "spaces"},
		},
		{
			input:    "UpperCase",
			expected: []string{"uppercase"},
		},
		{
			input:    "lowercase",
			expected: []string{"lowercase"},
		},
		{
			input:    "         ",
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := cleanInput(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// Helper function to compare slices (order doesn't matter)
func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	ma := make(map[string]bool)
	mb := make(map[string]bool)

	for _, val := range a {
		ma[val] = true
	}
	for _, val := range b {
		mb[val] = true
	}

	return reflect.DeepEqual(ma, mb)
}

func TestCleanInputUnordered(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This is a test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "    Extra    Spaces    ",
			expected: []string{"extra", "spaces"},
		},
		{
			input:    "UpperCase",
			expected: []string{"uppercase"},
		},
		{
			input:    "lowercase",
			expected: []string{"lowercase"},
		},
		{
			input:    "         ",
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := cleanInput(tc.input)
			if !slicesEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
