package chunk

import (
	"math"
)

/*
*
Creates a slice of elements split into groups the length of size.
If slice can't be split evenly, the final chunk will be the remaining elements.
*/
func Chunk[T interface{}](input []T, size int) [][]T {
	if size == 0 {
		return make([][]T, 0) // in js _.chunk([1,2,3], 0) === [], but the compiler won't let us here
	}
	numberOfChunks := int(math.Ceil(float64(len(input)) / float64(size)))
	result := make([][]T, numberOfChunks)
	chunkIndex := 0
	indexOfCurrentChunk := 0
	for i := 0; i < len(input); i++ {
		if result[chunkIndex] == nil {
			result[chunkIndex] = make([]T, size)
		} else if indexOfCurrentChunk == size {
			indexOfCurrentChunk = 0
			chunkIndex++
			result[chunkIndex] = make([]T, size)
		}
		result[chunkIndex][i%size] = input[i]
		indexOfCurrentChunk++
	}
	// slice final index to it's correct length (i.e the remainder)
	remainder := len(input) % size
	if remainder > 0 {
		result[chunkIndex] = result[chunkIndex][:len(input)%size]
	}
	return result
}
