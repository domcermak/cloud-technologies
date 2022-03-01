package main

const (
	Gas      = "gas"
	Diesel   = "diesel"
	Lpg      = "lpg"
	Electric = "electric"
)

var (
	EnergyTypes = []EnergyType{Gas, Diesel, Lpg, Electric}
)

type EnergyType string
