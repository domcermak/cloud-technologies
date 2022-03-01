package main

type Station struct {
	Entity
	energyType EnergyType
}

type StationConfig struct {
	EntityConfig
	energyType EnergyType
}

func NewStationConfig(cap uint32, waitDurationRange DurationRange, energyType EnergyType, next EntityInterface) *StationConfig {
	return &StationConfig{
		EntityConfig: *NewEntityConfig(cap, waitDurationRange, next),
		energyType:   energyType,
	}
}

func NewStation(config StationConfig) *Station {
	return &Station{
		Entity:     *NewEntity(config.EntityConfig),
		energyType: config.energyType,
	}
}

func (s *Station) Open() {
	go s.Entity.Open(
		func(car *Car) {
			car.ExitStationQueueAndStartPumping()
		},
		func(car *Car) {
			car.StopPumpingAndEnterCashMachineQueue()
		},
	)
}

func (s *Station) InstanceCount(count *uint32) {
	s.activeEntityInstances = count
}
