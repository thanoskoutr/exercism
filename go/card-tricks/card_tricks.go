package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// a the given index existed in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index >= len(slice) || index < 0 {
		return 0, false
	} else {
		return slice[index], true
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range it is be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length <= 0 {
		return nil
	} else {
		slice := make([]int, length)
		for i := 0; i < length; i++ {
			slice[i] = value
		}
		return slice
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= len(slice) || index < 0 {
		return slice
	} else {
		sliceNew := slice[:index]
		sliceNew = append(sliceNew, slice[index+1:]...)
		// for _, val := range slice[index+1:] {
		// 	sliceNew = append(sliceNew, val)
		// }
		return sliceNew
	}
}
