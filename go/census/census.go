// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	// panic("Please implement NewResident.")
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	// panic("Please implement HasRequiredInfo.")
	if r.Name != "" && r.Address["street"] != "" {
		return true
	}
	return false
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	// panic("Please implement Delete.")

	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	// panic("Please implement Count.")
	count := 0
	for _, value := range residents{
		if value.HasRequiredInfo() == true{
			count ++
		}
	}
	return count
}
