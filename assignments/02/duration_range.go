package main

import (
	"math/rand"
	"time"
)

type DurationRange struct {
	min time.Duration
	max time.Duration
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

	return time.Duration(rand.Float64()*(float64(r.max-r.min)) + float64(r.min))
}
