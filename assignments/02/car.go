package main

type Car struct {
	energyType                                           EnergyType
	stationQueue, station, cashMachineQueue, cashMachine StopWatch
}

func (c Car) EnergyType() EnergyType {
	return c.energyType
}

func (c *Car) EnterStationQueue() {
	c.stationQueue.Start()
}

func (c *Car) ExitStationQueueAndStartPumping() {
	c.stationQueue.Stop()
	c.station.Start()
}

func (c *Car) StopPumpingAndEnterCashMachineQueue() {
	c.station.Stop()
	c.cashMachineQueue.Start()
}

func (c *Car) ExitCashMachineQueueAndPay() {
	c.cashMachineQueue.Stop()
	c.cashMachine.Start()
}

func (c *Car) FinishPaymentAndLeave() {
	c.cashMachine.Stop()
}
