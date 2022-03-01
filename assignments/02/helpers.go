package main

import (
	"sync"
	"time"
)

func OpenAndAcceptCars(cashMachine *CashMachine, stations []*Station, carsToComePerStation uint) {
	cashMachine.Open()
	for _, station := range stations {
		station.Open()
	}

	time.Sleep(time.Second) // a delay to open all stations

	wg := &sync.WaitGroup{}
	wg.Add(len(stations))
	for _, station := range stations {
		go sendCars(station, carsToComePerStation, wg)
	}
	wg.Wait()

	for _, station := range stations {
		station.Close()
	}
}

func sendCars(station *Station, carsToComePerStation uint, wg *sync.WaitGroup) {
	for i := uint(0); i < carsToComePerStation; i++ {
		car := Car{energyType: station.energyType}
		car.EnterStationQueue()

		station.AcceptCar(car)
	}
	wg.Done()
}
