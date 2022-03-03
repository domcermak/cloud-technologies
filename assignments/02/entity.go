package main

import (
	"time"
)

type EntityInterface interface {
	Open()
	WaitUntilAvailable()
	ProcessCar()
	Close()
	String() string
}

type Entity struct {
	capacity                uint32
	waitDurationRange       DurationRange
	announcedAvailableSpots chan interface{}
}

type EntityConfig struct {
	capacity          uint32
	waitDurationRange DurationRange
}

func NewEntityConfig(cap uint32, waitDurationRange DurationRange) *EntityConfig {
	return &EntityConfig{
		capacity:          cap,
		waitDurationRange: waitDurationRange,
	}
}

func NewEntity(config EntityConfig) *Entity {
	return &Entity{
		capacity:                config.capacity,
		waitDurationRange:       config.waitDurationRange,
		announcedAvailableSpots: make(chan interface{}, config.capacity),
	}
}

func (e *Entity) Open() {
	for i := uint32(0); i < e.capacity; i++ {
		e.announcedAvailableSpots <- 0
	}
}

func (e *Entity) WaitUntilAvailable() {
	<-e.announcedAvailableSpots
}

func (e *Entity) ProcessCar() {
	time.Sleep(e.waitDuration())
	e.announcedAvailableSpots <- 0
}

func (e *Entity) Close() {
	close(e.announcedAvailableSpots)
}

func (e Entity) waitDuration() time.Duration {
	return e.waitDurationRange.RandomDuration()
}
