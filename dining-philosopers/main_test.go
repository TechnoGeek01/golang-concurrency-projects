package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != len(philosophers) {
			t.Errorf("incorrect length of slice; expected %d but got %d", len(philosophers), len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		orderFinished = []string{}

		eatTime = e.delay
		thinkTime = e.delay
		sleepTime = e.delay

		dine()
		if len(orderFinished) != len(philosophers) {
			t.Errorf("%s: incorrect length of slice; expected %d but got %d", e.name, len(philosophers), len(orderFinished))
		}
	}
}
