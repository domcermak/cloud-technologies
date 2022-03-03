package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

type Stats struct {
	TotalDuration       map[string]string `json:"total_duration"`
	MeanDuration        map[string]string `json:"mean_duration"`
	Sigma2Duration      map[string]string `json:"sigma_2_duration"`
	RunCount            map[string]int64  `json:"run_count"`
	CarCount            int64             `json:"car_count"`
	totalDuration       map[string]time.Duration
	meanDuration        map[string]float64
	sigma2Duration      map[string]float64
	partialMeasurements map[string][]time.Duration
	mutex               *sync.Mutex
}

func NewStats() *Stats {
	return &Stats{
		TotalDuration:       make(map[string]string),
		MeanDuration:        make(map[string]string),
		Sigma2Duration:      make(map[string]string),
		totalDuration:       make(map[string]time.Duration),
		meanDuration:        make(map[string]float64),
		sigma2Duration:      make(map[string]float64),
		RunCount:            make(map[string]int64),
		partialMeasurements: make(map[string][]time.Duration),
		mutex:               &sync.Mutex{},
	}
}

func (s *Stats) Measure(label string, fn func()) {
	start := time.Now()
	fn()
	end := time.Now()
	duration := end.Sub(start)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.RunCount[label]++
	s.totalDuration[label] += duration
	s.partialMeasurements[label] = append(s.partialMeasurements[label], duration)
}

func (s Stats) GetCarCount() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.CarCount
}

func (s *Stats) Done() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.CarCount++
}

func (s *Stats) Process() error {
	if s.CarCount == 0 {
		return errors.New("no cars were measured yet")
	}

	for label, total := range s.totalDuration {
		meanDuration := float64(total) / float64(s.RunCount[label])

		s.TotalDuration[label] = fmt.Sprintf("%v ms", total.Milliseconds())
		s.Sigma2Duration[label] = fmt.Sprintf(
			"%.2f",
			variance(s.partialMeasurements[label], meanDuration)/float64(time.Millisecond),
		)
		s.MeanDuration[label] = fmt.Sprintf("%.2f ms", meanDuration/float64(time.Millisecond))
	}

	return nil
}

func (s Stats) String() string {
	js, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Stats:\n%s", string(js))
}

func variance(measurements []time.Duration, mean float64) float64 {
	if len(measurements) == 0 {
		return 0.
	}

	sigma2 := float64(0.)
	for _, measurement := range measurements {
		sigma2 += math.Pow(float64(measurement)-mean, 2)
	}

	return sigma2 / float64(len(measurements))
}
