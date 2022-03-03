package main

type StationsConfig map[EnergyType]StationConfig

type Stations []*Station

func NewStationsConfig(gas, diesel, lpg, electric StationConfig) *StationsConfig {
	return &StationsConfig{
		Gas:      gas,
		Diesel:   diesel,
		Lpg:      lpg,
		Electric: electric,
	}
}

func DefaultStationsConfig() *StationsConfig {
	return NewStationsConfig(
		*NewStationConfig(4, *NewDurationRange(1*DurationUnit, 5*DurationUnit), Gas),
		*NewStationConfig(4, *NewDurationRange(1*DurationUnit, 5*DurationUnit), Diesel),
		*NewStationConfig(1, *NewDurationRange(1*DurationUnit, 5*DurationUnit), Lpg),
		*NewStationConfig(8, *NewDurationRange(3*DurationUnit, 10*DurationUnit), Electric),
	)
}

func InitStations(stationsConfig StationsConfig) Stations {
	stations := make([]*Station, len(stationsConfig))

	i := 0
	for _, stationConfig := range stationsConfig {
		func() {
			defer func() { i++ }()

			station := NewStation(stationConfig)
			stations[i] = station
		}()
	}

	return stations
}

func (s Stations) OpenAll() {
	for _, station := range s {
		station.Open()
	}
}

func (s Stations) CloseAll() {
	for _, station := range s {
		station.Close()
	}
}
