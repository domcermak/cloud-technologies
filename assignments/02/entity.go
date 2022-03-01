package main

import (
	"sync"
	"time"
)

type EntityInterface interface {
	AcceptCar(Car)
	Open()
	Close()
}

type Entity struct {
	capacity              uint32
	queue                 chan Car
	waitDurationRange     DurationRange
	quit                  chan interface{}
	next                  EntityInterface
	mutex                 *sync.Mutex
	activeEntityInstances *uint32
}

type EntityConfig struct {
	capacity          uint32
	waitDurationRange DurationRange
	next              EntityInterface
}

func NewEntityConfig(cap uint32, waitDurationRange DurationRange, next EntityInterface) *EntityConfig {
	return &EntityConfig{
		capacity:          cap,
		waitDurationRange: waitDurationRange,
		next:              next,
	}
}

func NewEntity(config EntityConfig) *Entity {
	active := config.capacity

	return &Entity{
		capacity:              config.capacity,
		queue:                 make(chan Car),
		quit:                  make(chan interface{}),
		waitDurationRange:     config.waitDurationRange,
		next:                  config.next,
		mutex:                 &sync.Mutex{},
		activeEntityInstances: &active,
	}
}

func (e Entity) Open(beforeProcessFn func(car *Car), afterProcessFn func(car *Car)) {
	for i := uint32(0); i < e.capacity; i++ {
		go func() {
			defer e.deactivate()

			for {
				select {
				case car := <-e.queue:
					beforeProcessFn(&car)
					time.Sleep(e.waitDuration())
					afterProcessFn(&car)

					if e.next != nil {
						e.next.AcceptCar(car)
					}
				case <-e.quit:
					return
				}
			}
		}()
	}
}

func (e Entity) Close() {
	for i := uint32(0); i < e.capacity; i++ {
		e.quit <- 0
	}
}

func (e Entity) AcceptCar(car Car) {
	e.queue <- car
}

func (e Entity) wasLastAlive() bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return *e.activeEntityInstances == 0
}

func (e Entity) deactivate() {
	func() {
		e.mutex.Lock()
		defer e.mutex.Unlock()

		if *e.activeEntityInstances == 0 {
			panic("No instance alive! There must be a concurrency issue!")
		}

		*e.activeEntityInstances -= 1
	}()

	if e.next != nil && e.wasLastAlive() { // allow just one thread to close next
		close(e.queue)

		// e.next.Close() must be outside of locks, otherwise it fails on deadlock
		e.next.Close()
	}
}

func (e Entity) waitDuration() time.Duration {
	return e.waitDurationRange.RandomDuration()
}
