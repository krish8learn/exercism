//Package weather this is how we should comment for package. 
package weather

//CurrentCondition this is how we should comment for variable declaration.
var CurrentCondition string
//CurrentLocation this is how we should comment for variable declaration.
var CurrentLocation string

//Forecast this is how we should comment for variable declaration.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
