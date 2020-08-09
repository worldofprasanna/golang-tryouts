package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("Invalid length")
	}
	for i, ch := range a {
		if rune(b[i]) != ch {
			distance = distance + 1
		}
	}
	return distance, nil
}
