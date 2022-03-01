package main

type CashMachine struct {
	Entity
}

type CashMachineConfig struct {
	EntityConfig
}

func NewCashMachineConfig(cap uint32, waitDurationRange DurationRange, next EntityInterface) *CashMachineConfig {
	return &CashMachineConfig{
		EntityConfig: *NewEntityConfig(cap, waitDurationRange, next),
	}
}

func NewCashMachine(config CashMachineConfig) *CashMachine {
	return &CashMachine{
		Entity: *NewEntity(config.EntityConfig),
	}
}

func DefaultCashMachineConfig(next EntityInterface) *CashMachineConfig {
	return NewCashMachineConfig(2, *NewDurationRange(DurationUnit/2, 2*DurationUnit), next)
}

func (cm *CashMachine) Open() {
	go cm.Entity.Open(
		func(car *Car) {
			car.ExitCashMachineQueueAndPay()
		},
		func(car *Car) {
			car.FinishPaymentAndLeave()
		},
	)
}
