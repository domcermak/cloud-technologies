package main

type CashMachine struct {
	Entity
}

type CashMachineConfig struct {
	EntityConfig
}

func NewCashMachineConfig(cap uint32, waitDurationRange DurationRange) *CashMachineConfig {
	return &CashMachineConfig{
		EntityConfig: *NewEntityConfig(cap, waitDurationRange),
	}
}

func NewCashMachine(config CashMachineConfig) *CashMachine {
	return &CashMachine{
		Entity: *NewEntity(config.EntityConfig),
	}
}

func DefaultCashMachineConfig() *CashMachineConfig {
	return NewCashMachineConfig(2, *NewDurationRange(DurationUnit/2, 2*DurationUnit))
}

func (_ CashMachine) String() string {
	return "cash_machine"
}
