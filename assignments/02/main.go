package main

import (
	"fmt"
	"time"
)

// Adjust following constants
const (
	// DurationUnit represents unit (second, millisecond, ...) of times
	// used as delays in stations and cash machines
	DurationUnit = time.Millisecond * 10

	// CarsToComePerStation represents number of cars sent to each station
	CarsToComePerStation = 1_000
)

func main() {
	stats := NewStats()
	cashMachineConfig := DefaultCashMachineConfig(stats)
	cashMachine := NewCashMachine(*cashMachineConfig)
	stationsConfig := DefaultStationsConfig(cashMachine)
	stations := InitStations(*stationsConfig)

	stats.Open()
	OpenAndAcceptCars(cashMachine, stations, CarsToComePerStation)
	stats.WaitUntilFinished()

	fmt.Println(stats)
}
