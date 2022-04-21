package purchase
import(
    "strings"
    "fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// panic("NeedsLicense not implemented")
    kind = strings.ToLower(kind)
    if kind == "car" || kind == "truck"{
        return true
    }
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// panic("ChooseVehicle not implemented")
    val := strings.Compare(option1,option2)
    if val == -1{
        return fmt.Sprintf("%s is clearly the better choice.", option1)
    }
	return fmt.Sprintf("%s is clearly the better choice.", option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// panic("CalculateResellPrice not implemented")
    var cost float64
    if age < 3{
        cost = (originalPrice * 80)/100
    }else if age >= 3 && age < 10 {
    	cost = (originalPrice * 70)/100
    }else if age >= 10{
    	cost = (originalPrice * 50)/100
    }
	return cost
}
