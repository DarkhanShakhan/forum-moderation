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

type Number interface {
	int64 | int32 | int | uint64 | uint32 | uint | float32 | float64
}

func IsPositiveNumber[T Number](nbr T) bool {
	return nbr > 0
}
