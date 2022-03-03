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
	fmt.Println("starting...")

	stats := NewStats()
	cashMachineConfig := DefaultCashMachineConfig()
	cashMachine := NewCashMachine(*cashMachineConfig)
	stationsConfig := DefaultStationsConfig()
	stations := InitStations(*stationsConfig)
	quitChan := make(chan interface{})

	fmt.Println("sending cars...")
	stats.Measure("whole_process", func() {
		SendCars(stations, cashMachine, stats, quitChan)
		stations.OpenAll()
		cashMachine.Open()

		fmt.Println("processing...")
		<-quitChan // wait until everything is processed

		fmt.Println("finishing...")
		stations.CloseAll()
		cashMachine.Close()
	})

	if err := stats.Process(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stats)
}
