package main

func getLast[T any](s []T) T {
	var lastElem T
	if len(s) == 0 {
		return lastElem
	}
	return s[len(s)-1]
}
