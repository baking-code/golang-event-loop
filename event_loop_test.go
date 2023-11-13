package main

import (
	"reflect"
	"testing"
	"time"
)

func TestLoop(t *testing.T) {
	var input = []func() string{
		func() string { time.Sleep(1 * time.Second); return "a" },
		func() string { time.Sleep(3 * time.Second); return "b" },
		func() string { time.Sleep(2 * time.Second); return "c" },
	}

	got := EventLoop[string](input)
	var expected = []string{"a", "c", "b"}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("EventLoop(); no good, got %s", got)
	}
}
