package meteorology

import (
	"fmt"
	"strconv"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	temperatureUnits := map[TemperatureUnit]string{
		Celsius: "°C", Fahrenheit: "°F",
	}

	return fmt.Sprintf(temperatureUnits[t])
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	tempValue := strconv.Itoa(t.degree)
	temUnit := t.unit.String()
	return fmt.Sprintf("%v %v", tempValue, temUnit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	speedUnit := map[SpeedUnit]string{
		KmPerHour: "km/h", MilesPerHour: "mph",
	}
	return fmt.Sprintf(speedUnit[s])
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	speedValue := strconv.Itoa(s.magnitude)
	speedUnit := s.unit.String()

	return fmt.Sprintf("%v %v", speedValue, speedUnit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string {
	loc := m.location
	temp := m.temperature.String()
	windDir := m.windDirection
	windSp := m.windSpeed.String()
	humid := strconv.Itoa(m.humidity) + "%"

	//"Athens: 21 °C, Wind N at 16 km/h, 63% Humidity"
	retString := fmt.Sprintf("%v: %v, Wind %v at %v, %v Humidity", loc, temp, windDir, windSp, humid)
	return retString
}
