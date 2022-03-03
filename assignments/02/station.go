package main

import "fmt"

type Station struct {
	Entity
	energyType EnergyType
}

type StationConfig struct {
	EntityConfig
	energyType EnergyType
}

func NewStationConfig(cap uint32, waitDurationRange DurationRange, energyType EnergyType) *StationConfig {
	return &StationConfig{
		EntityConfig: *NewEntityConfig(cap, waitDurationRange),
		energyType:   energyType,
	}
}

func NewStation(config StationConfig) *Station {
	return &Station{
		Entity:     *NewEntity(config.EntityConfig),
		energyType: config.energyType,
	}
}

func (s Station) String() string {
	return fmt.Sprintf("%s_station", s.energyType)
}
