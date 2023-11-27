package difference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	var input = []string{"a", "b", "c", "d"}

	var compare = []string{"b", "d"}

	got := Difference[string](input, compare)
	var expected = []string{"a", "c"}
	assert.EqualValues(t, expected, got)
}
func TestDifferenceEmpty(t *testing.T) {
	var input = []string{}

	var compare = []string{}

	got := Difference[string](input, compare)
	var expected = []string{}
	assert.EqualValues(t, expected, got)
}
func TestDifferenceSame(t *testing.T) {
	var input = []string{"a", "b"}

	var compare = []string{"a", "b"}

	got := Difference[string](input, compare)
	var expected = []string{}
	assert.EqualValues(t, expected, got)
}

func TestDifferenceLarge(t *testing.T) {
	var input = []string{"a", "b"}

	var compare = []string{"a", "b", "c"}

	got := Difference[string](input, compare)
	var expected = []string{}
	assert.EqualValues(t, expected, got)
}
