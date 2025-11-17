package main

func getLast[T any](s []T) T {
	var myZero T
	length := len(s)
	if length == 0 {
		return myZero
	}
	return s[length - 1]
}


// next step: https://www.boot.dev/lessons/c8999752-768a-401b-b881-602929927699