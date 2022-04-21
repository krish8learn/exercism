package partyrobot
import(
    "fmt"
    "strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome function")
    ret := fmt.Sprintf("Welcome to my party, %s!",name)
    return ret
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")
    ret1 := fmt.Sprintf("Happy birthday %s!", name)
    ret2 := fmt.Sprintf(" You are now %v years old!", age)
    return ret1 + ret2 
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// panic("Please implement the AssignTable function")
    ret1 := fmt.Sprintf("Welcome to my party, %s!",name)
    tableString := strconv.Itoa(table)
    for len(tableString) != 3{
        tableString = "0"+tableString
    }
    ret2 := fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %.1f meters from here.",tableString,direction,distance)
    ret3 := fmt.Sprintf("You will be sitting next to %s.", neighbor)
    ret := ret1 + "\n" + ret2 + "\n" + ret3
    return ret
}
