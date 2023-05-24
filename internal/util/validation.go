package util

import "strings"

func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsEmptySlice[T comparable](s []T) bool {
	return len(s) == 0
}

func IsUniqueSlice[T comparable](s []T) bool {
	unique := map[T]bool{}
	for _, v := range s {
		unique[v] = true
	}
	return len(unique) == len(s)
}
