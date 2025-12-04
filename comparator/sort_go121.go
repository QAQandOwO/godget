//go:build go1.21
// +build go1.21

package comparator

import "slices"

// SortSlice sorts the slice using the comparator. Not stable.
func (c *Comparator[T]) SortSlice(slice []T) {
	slices.SortFunc(slice, c.Compare)
}

// SortSliceStable sorts the slice stably using the comparator.
func (c *Comparator[T]) SortSliceStable(slice []T) {
	slices.SortStableFunc(slice, c.Compare)
}

// SliceIsSorted checks if the slice is sorted according to the comparator.
func (c *Comparator[T]) SliceIsSorted(slice []T) bool {
	return slices.IsSortedFunc(slice, c.Compare)
}
