package main

import "fmt"

/*
Execute array of asynchronous functions and return array in order of response.
Mimics behaviour of Javascript's Event loop
*/
func EventLoop[T interface{}](inputFunctions []func() T) []T {
	var toReturn []T
	var errors []error

	mainChannel := make(chan T, 1)
	done := make(chan bool)

	go func() {
		for {
			value, isMore := <-mainChannel
			if isMore {
				toReturn = append(toReturn, value)
				if len(toReturn)+len(errors) == cap(inputFunctions) {
					close(mainChannel)
				}
			} else {
				// mainChannel closed, trigger done channel
				done <- true
				return
			}
		}
	}()

	for index, inputFunction := range inputFunctions {
		executor := inputFunction
		i := index
		go func() {
			defer func() {
				if r := recover(); r != nil {
					// catch any unhandled errors, print and continue execution. Like wrapping in a try/catch
					errors = append(errors, fmt.Errorf("Received error from input function %d: %s", i, r))
				}
			}()
			mainChannel <- executor()
		}()
	}
	// wait for done channel to receive signal
	<-done
	if len(errors) > 0 {
		fmt.Printf("%+v\n", errors)
	}
	return toReturn
}

/*
Implements javascripts Promise.all by evaluating array of async
operations in parallel, but returning their values in their given
order
*/
func PromiseAll[T interface{}](inputFunctions []func() T) []T {
	var toReturn = make([]T, len(inputFunctions), len(inputFunctions))
	type ValueIndexTuple struct {
		value T
		index int
	}
	mainChannel := make(chan ValueIndexTuple, 1)
	done := make(chan bool)

	go func() {
		times := 0
		for {
			tuple, isMore := <-mainChannel
			if isMore {

				toReturn[tuple.index] = tuple.value
				times++
				if times == cap(inputFunctions) {
					close(mainChannel)
				}
			} else {
				// mainChannel closed, trigger done channel
				done <- true
				return
			}
		}
	}()

	for index, inputFunction := range inputFunctions {
		executor := inputFunction
		i := index
		go func() {
			mainChannel <- ValueIndexTuple{value: executor(), index: i}
		}()
	}
	// wait for done channel to receive signal
	<-done
	return toReturn
}
