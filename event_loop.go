package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func EventLoop[T interface{}](inputFunctions []func() T) []T {
	var toReturn []T

	mainChannel := make(chan T, 1)
	done := make(chan bool)

	go func() {
		for {
			value, isMore := <-mainChannel
			if isMore {
				toReturn = append(toReturn, value)
				if len(toReturn) == cap(inputFunctions) {
					close(mainChannel)
				}
			} else {
				// mainChannel closed, trigger done channel
				done <- true
				return
			}
		}
	}()

	for _, inputFunction := range inputFunctions {
		executor := inputFunction
		go func() {
			mainChannel <- executor()
		}()
	}
	// wait for done channel to receive signal
	<-done
	return toReturn
}
