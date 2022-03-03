package main

import (
	"math/rand"
	"time"
)

type DurationRange struct {
	min, max time.Duration
}

func NewDurationRange(min, max time.Duration) *DurationRange {
	return &DurationRange{
		min: min,
		max: max,
	}
}

func (r DurationRange) RandomDuration() time.Duration {
	if r.min == r.max {
		return r.min
	}

	return time.Duration(rand.Int63n(int64(r.max-r.min)) + int64(r.min))
}
