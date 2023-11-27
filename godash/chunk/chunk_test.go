package chunk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	var input = []string{"a", "b", "c", "d", "e"}

	got := Chunk[string](input, 2)
	var expected = [][]string{{"a", "b"}, {"c", "d"}, {"e"}}
	assert.EqualValues(t, expected, got)
}
func TestPerfectChunk(t *testing.T) {
	var input = []string{"a", "b", "c", "d", "e", "f"}

	got := Chunk[string](input, 2)
	var expected = [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	assert.EqualValues(t, expected, got)
}

func TestOneChunk(t *testing.T) {
	var input = []string{"a", "b", "c"}

	got := Chunk[string](input, 1)
	var expected = [][]string{{"a"}, {"b"}, {"c"}}
	assert.EqualValues(t, expected, got)
}

func TestTooManyChunks(t *testing.T) {
	var input = []string{"a", "b", "c"}

	got := Chunk[string](input, 4)
	var expected = [][]string{{"a", "b", "c"}}
	assert.EqualValues(t, expected, got)
}

func TestNoChunks(t *testing.T) {
	var input = []string{"a", "b", "c", "d", "e"}

	got := Chunk[string](input, 0)
	var expected = [][]string{}
	assert.EqualValues(t, expected, got)
}
