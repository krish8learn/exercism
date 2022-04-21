package gross


// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// panic("Please implement the Units() function")
	grossUnit := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return grossUnit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	// panic("Please implement the NewBill() function")
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the AddItem() function")
	if _, exists := units[unit]; exists == false {
		//unit is not in units map
		return false
	}
	//unit is in units map
	if _, exists := bill[item]; exists == true {
		//item already exists in the bill
		bill[item] += units[unit]
		return true
	}
	//item not in the bill
	bill[item] = units[unit]
	// fmt.Println("-->", bill, "-->", units)
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the RemoveItem() function")
	if _, exists := bill[item]; exists == false {
		//item is not in bill map
		return false
	}

	//item is in the bill

	if _, exists := units[unit]; exists == false {
		//unit is not in units map
		return false
	}

	//unit is in the unit map
	testValue :=bill[item]-units[unit]
	if testValue < 0 {
		return false
	} else if testValue == 0 {
		delete(bill, item)
		return true
	}

    bill[item] -= units[unit]
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	// panic("Please implement the GetItem() function")
	if value, exists := bill[item]; exists == true {
		return value, true
	}
	return 0, false
}
