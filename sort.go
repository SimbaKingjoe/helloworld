// sort.go
package mypackage

import (
	"sort"
)

// SortInts sorts a slice of integers in ascending order.
func SortInts(slice []int) {
	sort.Ints(slice)
}

// SortStrings sorts a slice of strings by length.
func SortStringsByLength(slice []string) {
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) < len(slice[j])
	})
}
