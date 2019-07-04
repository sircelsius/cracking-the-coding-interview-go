package arrays_and_strings

import "strings"

// 1.1 Is Unique: implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?
func IsUnique(source string) bool {
	var seen = ""
	for _, s := range source {
		if strings.Contains(seen, string(s)) {
			return false
		}
		seen += string(s)
	}
	return true
}