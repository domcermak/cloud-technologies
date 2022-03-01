package main

import (
	"errors"
	"fmt"
	"time"
)

type StopWatch struct {
	startTime, stopTime time.Time
}

func (sw *StopWatch) Start() {
	sw.startTime = time.Now()
}

func (sw *StopWatch) Stop() {
	sw.stopTime = time.Now()
}

func (sw StopWatch) Duration() (time.Duration, error) {
	if sw.startTime.Equal(time.Time{}) ||
		sw.stopTime.Equal(time.Time{}) ||
		sw.startTime.After(sw.stopTime) {
		return time.Nanosecond, errors.New(
			fmt.Sprintf(
				"Invalid time range: %v; %v",
				sw.startTime.Format(time.RFC3339), sw.stopTime.Format(time.RFC3339)),
		)
	}

	return sw.stopTime.Sub(sw.startTime), nil
}
