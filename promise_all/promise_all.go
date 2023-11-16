package promise

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
