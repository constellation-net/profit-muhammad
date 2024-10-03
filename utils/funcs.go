package utils

import (
	"math/rand"
)

// RandSliceItem returns a random item from the given slice.
func RandSliceItem(a []string) string {
	if len(a) == 0 {
		return ""
	}

	i := rand.Intn(len(a))
	return a[i]
}

// SliceContains determines if a slice of strings contains an element exactly equal to a given string.
func SliceContains(s []string, v string) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}

	return false
}
