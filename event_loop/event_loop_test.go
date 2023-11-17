package eventloop

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventLoop(t *testing.T) {
	var input = []func() string{
		func() string { return "a" },
		func() string { return "b" },
		func() string { return "c" },
		func() string { return "d" },
		func() string { return "e" },
	}

	got := EventLoop[string](input)
	var expected = []string{"a", "b", "c", "d", "e"}
	assert.ElementsMatch(t, expected, got)
}

func TestEventLoopOutOfOrder(t *testing.T) {
	var input = []func() string{
		func() string { time.Sleep(3 * time.Second); return "a" },
		func() string { time.Sleep(2 * time.Second); return "b" },
		func() string { time.Sleep(5 * time.Second); return "c" },
		func() string { time.Sleep(1 * time.Second); return "d" },
		func() string { time.Sleep(4 * time.Second); return "e" },
	}

	got := EventLoop[string](input)
	var expected = []string{"d", "b", "a", "e", "c"}
	assert.EqualValues(t, expected, got)
}

func TestEventLoopParallel(t *testing.T) {
	timeStart := time.Now()
	var input = []func() string{
		func() string { time.Sleep(4 * time.Second); return "a" },
		func() string { time.Sleep(1 * time.Second); return "b" },
		func() string { time.Sleep(5 * time.Second); return "c" },
	}

	got := EventLoop[string](input)
	elapsed := time.Since((timeStart))
	var expected = []string{"b", "a", "c"}
	assert.EqualValues(t, expected, got)
	// verify that we're running in parallel
	assert.LessOrEqual(t, elapsed.Seconds(), 7.0)
}

func TestEventLoopErrorHandling(t *testing.T) {
	var input = []func() string{
		func() string { time.Sleep(3 * time.Second); return "a" },
		func() string { time.Sleep(1 * time.Second); return "b" },
		func() string { time.Sleep(2 * time.Second); panic("help!") },
	}

	got := EventLoop[string](input)
	var expected = []string{"b", "a"}
	assert.EqualValues(t, expected, got)

}
