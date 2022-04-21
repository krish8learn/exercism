package speed

// TODO: define the 'Car' type struct
type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	//panic("Please implement the NewCar function")
    return Car{
        battery:100,
        batteryDrain: batteryDrain,
        speed: speed,
        distance:0,
    }
}

// TODO: define the 'Track' type struct
type Track struct{
    distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	//panic("Please implement the NewTrack function")
    return Track{
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive on more //time,
// the car will not move.
func Drive(car Car) Car {
	//panic("Please implement the Drive function")
    if car.battery < car.batteryDrain{
        return car
    }
    car.battery = car.battery-car.batteryDrain
    car.distance += car.speed
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	//panic("Please implement the CanFinish function")
    if car.battery < car.batteryDrain {
        return false
    }
    maxDistance := (car.battery/car.batteryDrain) * car.speed
    if maxDistance >= track.distance{
        return true
    }
	return false
}
