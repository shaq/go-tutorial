package main

import "fmt"

const (
	uint16Max   float64 = 65535
	kmhMultiple float64 = 1.60934
)

type car struct {
	throttle      uint16
	brake         uint16
	steeringWheel int16
	topSpeedInKMH float64
}

// This is a method, not a function!
func (c car) kmh() float64 {
	return float64(c.throttle) * (c.topSpeedInKMH / uint16Max)
}

// This is a value-receiver method
func (c car) mph() float64 {
	return float64(c.throttle) * (c.topSpeedInKMH / uint16Max / kmhMultiple)
}

// This is a pointer-receiver method.
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedInKMH = newSpeed
}

// function that takes a car, modifies the speed and returns
// that same car modified. Clearly not as efficient as using a
// pointer receiver.
func newerTopSpeed(c car, newSpeed float64) car {
	c.topSpeedInKMH = newSpeed
	return c
}

func main() {
	aCar := car{throttle: 65000, brake: 0, steeringWheel: 65, topSpeedInKMH: 225.0}

	fmt.Println(aCar.throttle)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	// aCar.newTopSpeed(300)
	aCar = newerTopSpeed(aCar, 300)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
