package main

import (
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	PanicAfter(0)
}

func TestPanic10Seconds(t *testing.T) {
	start := time.Now()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
		d := time.Since(start)
		if d.Seconds() < 9 || d.Seconds() > 11 {
			t.Errorf("The code did not panic in expected time")
		}
	}()

	PanicAfter(10)
}
