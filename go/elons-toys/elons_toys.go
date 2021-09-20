package elon

import "fmt"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// Track implements a race track.
type Track struct {
	distance int
}

// CreateCar creates a new car with given specifications.
func CreateCar(speed, batteryDrain int) *Car {
	tempCar := Car{speed: speed, batteryDrain: batteryDrain, battery: 100}
	return &tempCar
}

// CreateTrack creates a new track with given distance.
func CreateTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time.
func (car *Car) Drive() {
	if car.batteryDrain <= car.battery {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// CanFinish checks if a car is able to finish a certain track.
func (car *Car) CanFinish(track Track) bool {
	remaining := (car.battery / car.batteryDrain) * car.speed
	if track.distance <= remaining {
		return true
	} else {
		return false
	}
}

// DisplayDistance displays the distance the car is driven.
func (car *Car) DisplayDistance() string {
	msg := fmt.Sprintf("Driven %v meters", car.distance)
	return msg
}

// DisplayBattery displays the battery level.
func (car *Car) DisplayBattery() string {
	msg := fmt.Sprintf("Battery at %v%%", car.battery)
	return msg
}
