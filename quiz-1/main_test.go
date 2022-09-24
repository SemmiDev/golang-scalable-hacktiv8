package main

import "testing"

func TestSort(t *testing.T) {
	var tests = []struct {
		name          string
		input, output []string
	}{
		{"test 1", []string{"apel", "mangga", "apel", "jeruk", "mangga", "jeruk", "apel", "mangga", "mangga"}, []string{"mangga 4", "apel 3", "jeruk 2"}},
		{"test 2", []string{"apel", "mangga", "jeruk", "apel", "jeruk", "apel", "mangga", "mangga", "mangga", "mangga"}, []string{"mangga 5", "apel 3", "jeruk 2"}},
		{"test 3", []string{"apel", "mangga", "jeruk", "apel", "jeruk", "apel", "mangga", "mangga", "mangga", "mangga", "mangga"}, []string{"mangga 6", "apel 3", "jeruk 2"}},
		{"test 4", []string{"apel", "apel", "apel", "apel", "apel", "apel", "mangga"}, []string{"apel 6", "mangga 1"}},
		{"test 5", []string{"apel", "apel", "mangga", "mangga"}, []string{"apel 2", "mangga 2"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			sortedFruitsWithCounts := SortFruits(test.input)
			if !equal(sortedFruitsWithCounts, test.output) {
				t.Errorf("got %v, want %v", sortedFruitsWithCounts, test.output)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
