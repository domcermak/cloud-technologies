package main

func SendCars(stations Stations, cashMachine *CashMachine, stats *Stats, quitChan chan<- interface{}) {
	for i := 0; i < CarsToComePerStation; i++ {
		for _, station := range stations {

			// this function represents a car
			go func(quitChan chan<- interface{}, stats *Stats, entities ...EntityInterface) {
				for _, entity := range entities {
					stats.Measure(entity.String()+"_wait_queue", func() {
						entity.WaitUntilAvailable()
					})
					stats.Measure(entity.String()+"_processing", func() {
						entity.ProcessCar()
					})
				}
				stats.Done()

				if stats.GetCarCount() == CarsToComePerStation*int64(len(stations)) {
					quitChan <- 0
				}
			}(quitChan, stats, station, cashMachine)
		}
	}
}
