package arrays_and_strings

import "testing"

func TestIsUnique(t *testing.T) {
	if !IsUnique("") {
		t.Errorf("Empty string should be considered unique")
	}

	if !IsUnique("abc") {
		t.Errorf("\"abc\" should be returned as unique")
	}

	if IsUnique("aabc") {
		t.Errorf("\"aabc\" should be returned as not unique")
	}
}
