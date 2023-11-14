package main

import (
	"reflect"
	"testing"
	"time"
)

func TestLoopOrder(t *testing.T) {
	var input = []func() string{
		func() string { time.Sleep(3 * time.Second); return "a" },
		func() string { time.Sleep(2 * time.Second); return "b" },
		func() string { time.Sleep(5 * time.Second); return "c" },
		func() string { time.Sleep(1 * time.Second); return "d" },
		func() string { time.Sleep(4 * time.Second); return "e" },
	}

	got := EventLoop[string](input)
	var expected = []string{"d", "b", "a", "e", "c"}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("EventLoop(); no good, got %s", got)
	}
}

func TestLoopParallel(t *testing.T) {
	timeStart := time.Now()
	var input = []func() string{
		func() string { time.Sleep(4 * time.Second); return "a" },
		func() string { time.Sleep(1 * time.Second); return "b" },
		func() string { time.Sleep(5 * time.Second); return "c" },
	}

	got := EventLoop[string](input)
	elapsed := time.Since((timeStart))
	var expected = []string{"b", "a", "c"}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("EventLoop(); no good, got %s", got)
	}
	// verify that we're running in parallel
	if elapsed.Seconds() > 7 {
		t.Errorf("EventLoop(); should have been run in parallel. Expected: between 5 and 6s; Actual: %d", elapsed)
	}
}

func TestLoopErrorHandling(t *testing.T) {
	var input = []func() string{
		func() string { time.Sleep(3 * time.Second); return "a" },
		func() string { time.Sleep(1 * time.Second); return "b" },
		func() string { time.Sleep(2 * time.Second); panic("help!") },
	}

	got := EventLoop[string](input)
	var expected = []string{"b", "a"}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("EventLoop(); no good, got %s", got)
	}
}
