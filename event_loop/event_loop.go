package event_loop

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
