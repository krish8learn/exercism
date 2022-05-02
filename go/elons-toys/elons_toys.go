package elon

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery > c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}

}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	ret := fmt.Sprintf("Driven %v meters", c.distance)
	return ret
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	ret := fmt.Sprintf("Battery at %v", c.battery)
	return ret + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
	if c.battery < c.batteryDrain {
		return false
	}
	maxDistance := (c.battery / c.batteryDrain) * c.speed
	if maxDistance >= trackDistance {
		return true
	}
	return false
}
