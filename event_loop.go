package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func EventLoop[T interface{}](inputFunctions []func() T) []T {
	var toReturn []T
	var wrap = func(mainChannel chan T, executor func() T, index int) {
		mainChannel <- executor()
		if index == cap(mainChannel) {
			close(mainChannel)
		}
	}
	mainChannel := make(chan T, 1)
	for idx, inputFunction := range inputFunctions {
		go wrap(mainChannel, inputFunction, idx)
	}
	for val := range mainChannel {
		toReturn = append(toReturn, val)
	}
	return toReturn
}
