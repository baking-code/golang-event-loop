package promise

import "sync"

/*
Implements javascripts Promise.all by evaluating array of async
operations in parallel, but returning their values in their given
order
*/
func PromiseAll[T interface{}](inputFunctions []func() T) []T {
	var toReturn = make([]T, len(inputFunctions))

	var wg sync.WaitGroup
	wg.Add(len(inputFunctions))
	for index, inputFunction := range inputFunctions {
		executor := inputFunction
		i := index
		go func() {
			defer wg.Done()
			toReturn[i] = executor()
		}()
	}
	wg.Wait()
	return toReturn
}
