package difference

import "slices"

/*
Creates an slice of input values not included in the other given slice using comparable.
The order and references of result values are determined by the first slice.
*/
func Difference[T comparable](input []T, compare []T) []T {
	result := make([]T, 0)

	for i := 0; i < len(input); i++ {
		element := input[i]
		if !slices.Contains(compare, element) {
			result = append(result, element)
		}
	}
	return result
}
