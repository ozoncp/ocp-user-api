package utils

import (
	"testing"
)

func TestRepeatFileContent(t *testing.T) {
	tests := []struct {
		description string
		path        string
		count       int
		expected    int
		err         bool
	}{
		{"FileNotExistsError", "1", 1, 0, true},
		{"FilePathEmptyError", "", 1, 0, true},
		{"CountNegativeError", "file_content.go", -1, 0, true},
		{"ZeroCount", "file_content.go", 0, 0, false},
		{"SomeCount", "file_content.go", 3, 3, false},
	}

	for _, item := range tests {
		actual, err := RepeatFileContent(item.path, item.count)

		if item.err {
			if err == nil {
				t.Errorf("%s: expected error", item.description)
			}
		} else {
			if actual == nil || len(actual) != item.expected {
				t.Errorf("%s: expected %v, but got %v", item.description, item.expected, actual)
			}
		}
	}
}
