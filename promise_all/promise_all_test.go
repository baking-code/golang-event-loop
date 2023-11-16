package promise

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPromiseAll(t *testing.T) {
	var input = []func() string{
		func() string { time.Sleep(3 * time.Second); return "a" },
		func() string { time.Sleep(2 * time.Second); return "b" },
		func() string { time.Sleep(5 * time.Second); return "c" },
		func() string { time.Sleep(1 * time.Second); return "d" },
		func() string { time.Sleep(4 * time.Second); return "e" },
	}

	got := PromiseAll[string](input)
	var expected = []string{"a", "b", "c", "d", "e"}
	assert.EqualValues(t, expected, got)
}
