package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"
)

type energyStatistics struct {
	Mean   float64 `json:"mean"`
	Sigma2 float64 `json:"sigma_2"`
}

type statistics struct {
	DurationInStationQueue     energyStatistics `json:"duration_in_station_queue"`
	DurationInStation          energyStatistics `json:"duration_in_station"`
	DurationInCashMachineQueue energyStatistics `json:"duration_in_cash_machine_queue"`
	DurationInCashMachine      energyStatistics `json:"duration_in_cash_machine"`
	CarCount                   int              `json:"car_count"`
}

type Stats struct {
	buffer map[EnergyType][]Car
	queue  chan Car
	quit   chan interface{}
}

func NewStats() *Stats {
	return &Stats{
		buffer: make(map[EnergyType][]Car),
		queue:  make(chan Car, 0),
		quit:   make(chan interface{}, 0),
	}
}

func (s Stats) AcceptCar(car Car) {
	s.queue <- car
}

func (s *Stats) Open() {
	go func() {
		i := 1
		for {
			car := <-s.queue

			mutex := sync.Mutex{}
			mutex.Lock()
			s.buffer[car.EnergyType()] = append(s.buffer[car.EnergyType()], car)

			fmt.Println(fmt.Sprintf("Car %d added {%s}", i, car.energyType))
			i++
			mutex.Unlock()
		}
	}()
}

func (s Stats) Close() {
	s.quit <- 0
}

func (s Stats) WaitUntilFinished() {
	<-s.quit
	close(s.quit)
}

func (s Stats) String() string {
	stats, err := s.process()
	if err != nil {
		return err.Error()
	}

	js, err := json.MarshalIndent(stats, "", "\t")
	if err != nil {
		return err.Error()
	}

	return string(js)
}

func (s Stats) process() (map[EnergyType]statistics, error) {
	mutex := sync.Mutex{}
	mutex.Lock()

	stats := make(map[EnergyType]statistics)
	for energyType, cars := range s.buffer {
		stat, err := statsForCars(cars)
		if err != nil {
			return nil, err
		}

		stats[energyType] = stat
	}

	mutex.Unlock()

	return stats, nil
}

func statsForCars(cars []Car) (statistics, error) {
	sqStat, err := energyStatsForCar(cars, func(car Car) (time.Duration, error) {
		return car.stationQueue.Duration()
	})
	if err != nil {
		return statistics{}, err
	}

	sStat, err := energyStatsForCar(cars, func(car Car) (time.Duration, error) {
		return car.station.Duration()
	})
	if err != nil {
		return statistics{}, err
	}

	cmqStat, err := energyStatsForCar(cars, func(car Car) (time.Duration, error) {
		return car.cashMachineQueue.Duration()
	})
	if err != nil {
		return statistics{}, err
	}

	cmStat, err := energyStatsForCar(cars, func(car Car) (time.Duration, error) {
		return car.cashMachine.Duration()
	})
	if err != nil {
		return statistics{}, err
	}

	return statistics{
		DurationInStationQueue:     sqStat,
		DurationInStation:          sStat,
		DurationInCashMachineQueue: cmqStat,
		DurationInCashMachine:      cmStat,
		CarCount:                   len(cars), // this is number of cars, that actually came
	}, nil
}

func energyStatsForCar(cars []Car, fn func(car Car) (time.Duration, error)) (energyStatistics, error) {
	mean, err := mean(cars, fn)
	if err != nil {
		return energyStatistics{}, err
	}
	sigma2, err := variance(cars, mean, fn)
	if err != nil {
		return energyStatistics{}, err
	}

	return energyStatistics{
		Mean:   mean,
		Sigma2: sigma2,
	}, nil
}

func mean(cars []Car, fn func(car Car) (time.Duration, error)) (float64, error) {
	mean := 0.

	for _, car := range cars {
		duration, err := fn(car)
		if err != nil {
			return 0., err
		}

		mean += float64(duration) / float64(DurationUnit)
	}

	return mean / float64(len(cars)), nil
}

func variance(cars []Car, mean float64, fn func(car Car) (time.Duration, error)) (float64, error) {
	sigma2 := 0.

	for _, car := range cars {
		duration, err := fn(car)
		if err != nil {
			return 0., err
		}
		sigma2 += math.Pow(float64(duration)/float64(DurationUnit)-mean, 2)
	}

	return sigma2 / float64(len(cars)), nil
}
