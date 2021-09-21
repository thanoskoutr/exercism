package hamming

import "errors"

func Distance(a, b string) (int, error) {
	// The Hamming distance is only defined for sequences of equal length
	if len(a) != len(b) {
		return 0, errors.New("")
	}
	// Calculate Hamming distance
	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
