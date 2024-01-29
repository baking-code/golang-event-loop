package eventloop

import (
	"fmt"
	"sync"
)

/*
Execute array of asynchronous functions and return array in order of response.
Mimics behaviour of Javascript's Event loop
*/
func EventLoop[T interface{}](inputFunctions []func() T) []T {
	var toReturn []T
	var errors []error

	var wg sync.WaitGroup
	wg.Add(len(inputFunctions))
	for index, inputFunction := range inputFunctions {
		executor := inputFunction
		i := index
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					// catch any unhandled errors, print and continue execution. Like wrapping in a try/catch
					errors = append(errors, fmt.Errorf("received error from input function %d: %s", i, r))
				}
			}()
			toReturn = append(toReturn, executor())
		}()
	}
	// wait for done channel to receive signal
	wg.Wait()
	if len(errors) > 0 {
		fmt.Printf("%+v\n", errors)
	}
	return toReturn
}
