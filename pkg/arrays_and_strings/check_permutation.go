package arrays_and_strings

import (
	"sort"
	"strings"
)

func CheckPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aSorted := strings.Split(a, "")
	sort.Strings(aSorted)
	bSorted := strings.Split(b, "")
	sort.Strings(bSorted)
	return strings.Join(aSorted, "") == strings.Join(bSorted, "")
}
