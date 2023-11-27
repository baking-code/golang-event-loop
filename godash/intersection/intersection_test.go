package intersection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	var input = []string{"a", "b", "c", "d"}

	var compare = []string{"b", "d"}

	got := Intersection[string](input, compare)
	var expected = []string{"b", "d"}
	assert.EqualValues(t, expected, got)
}
func TestIntersectionEmpty(t *testing.T) {
	var input = []string{}

	var compare = []string{}

	got := Intersection[string](input, compare)
	var expected = []string{}
	assert.EqualValues(t, expected, got)
}
func TestIntersectionSame(t *testing.T) {
	var input = []string{"a", "b"}

	var compare = []string{"a", "b"}

	got := Intersection[string](input, compare)
	var expected = []string{"a", "b"}
	assert.EqualValues(t, expected, got)
}

func TestIntersectionOpposite(t *testing.T) {
	var input = []string{"a", "b"}

	var compare = []string{"a", "b", "c"}

	got := Intersection[string](input, compare)
	var expected = []string{"a", "b"}
	assert.EqualValues(t, expected, got)
}
