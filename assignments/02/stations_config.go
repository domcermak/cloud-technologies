package main

type StationsConfig map[EnergyType]StationConfig

func NewStationsConfig(gas, diesel, lpg, electric StationConfig) *StationsConfig {
	return &StationsConfig{
		Gas:      gas,
		Diesel:   diesel,
		Lpg:      lpg,
		Electric: electric,
	}
}

func DefaultStationsConfig(next EntityInterface) *StationsConfig {
	return NewStationsConfig(
		*NewStationConfig(4, *NewDurationRange(1*DurationUnit, 5*DurationUnit), Gas, next),
		*NewStationConfig(4, *NewDurationRange(1*DurationUnit, 5*DurationUnit), Diesel, next),
		*NewStationConfig(1, *NewDurationRange(1*DurationUnit, 5*DurationUnit), Lpg, next),
		*NewStationConfig(8, *NewDurationRange(3*DurationUnit, 10*DurationUnit), Electric, next),
	)
}

func InitStations(stationsConfig StationsConfig) []*Station {
	stations := make([]*Station, len(stationsConfig))
	allStationCount := uint32(0)

	i := 0
	for _, stationConfig := range stationsConfig {
		func() {
			defer func() { i++ }()

			station := NewStation(stationConfig)

			stations[i] = station
			stations[i].InstanceCount(&allStationCount)

			allStationCount += station.capacity
		}()
	}

	return stations
}
