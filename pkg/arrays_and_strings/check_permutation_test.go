package arrays_and_strings

import "testing"

// 1.2 Check Permutation: Given two strings, write a method to decide if one is a permutation of
// the other.
func TestCheckPermutation(t *testing.T) {
	if !CheckPermutation("abc", "cba") {
		t.Errorf("abc and cba should be a permutation!")
	}

	if CheckPermutation("cba", "cbaa") {
		t.Errorf("cba and cbaa are not permutations!")
	}
}
