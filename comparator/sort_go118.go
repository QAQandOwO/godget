//go:build !go1.21
// +build !go1.21

package comparator

import "sort"

// SortSlice sorts the slice using the comparator. Not stable.
func (c *Comparator[T]) SortSlice(slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return c.Less(slice[i], slice[j])
	})
}

// SortSliceStable sorts the slice stably using the comparator.
func (c *Comparator[T]) SortSliceStable(slice []T) {
	sort.SliceStable(slice, func(i, j int) bool {
		return c.Less(slice[i], slice[j])
	})
}

// SliceIsSorted checks if the slice is sorted according to the comparator.
func (c *Comparator[T]) SliceIsSorted(slice []T) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return c.Less(slice[i], slice[j])
	})
}
