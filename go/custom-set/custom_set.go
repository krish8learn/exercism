package stringset

import (
	"fmt"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]string

func New() Set {
	// panic("Please implement the New function")
	return make(Set)
}

func NewFromSlice(l []string) Set {
	// panic("Please implement the NewFromSlice function")
	var set = make(Set)
	for _, value := range l {
		set[value] = value
	}
	return set
}

func (s Set) String() string {
	// panic("Please implement the String function")
	if len(s) == 0 {
		return "{}"
	}

	index := 0
	stringValue := "{"
	for _, value := range s {
		if index != len(s)-1 {
			stringValue += fmt.Sprintf(`"%s", `, value)
		} else {
			stringValue += fmt.Sprintf(`"%s"`, value)
		}
		index++
	}

	stringValue += "}"

	return stringValue
}

func (s Set) IsEmpty() bool {
	// panic("Please implement the IsEmpty function")
	if len(s) == 0 {
		return true
	}
	return false
}

func (s Set) Has(elem string) bool {
	// panic("Please implement the Has function")
	if _, present := s[elem]; !present {
		return false
	}
	return true
}

func (s Set) Add(elem string) {
	// panic("Please implement the Add function")
	s[elem] = elem
}

func Subset(s1, s2 Set) bool {
	// panic("Please implement the Subset function")
	if s1.IsEmpty() && s2.IsEmpty() {
		return true
	} else if s1.IsEmpty() && !s2.IsEmpty() {
		return true
	}

	for key, _ := range s1 {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	// panic("Please implement the Disjoint function")
	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}

	for key, _ := range s1 {
		if s2.Has(key) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	// panic("Please implement the Equal function")
	if len(s1) != len(s2) {
		return false
	}
	for key, _ := range s1 {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	// panic("Please implement the Intersection function")
	intersectionSet := New()

	for key := range s1 {
		if s2.Has(key) {
			intersectionSet.Add(key)
		}
	}

	return intersectionSet
}

func Difference(s1, s2 Set) Set {
	// panic("Please implement the Difference function")
	differenceSet := New()

	for key := range s1 {
		if !s2.Has(key) {
			differenceSet.Add(key)
		}
	}

	return differenceSet

}

func Union(s1, s2 Set) Set {
	// panic("Please implement the Union function")
	unionSet := New()

	for key, _ := range s1 {
		unionSet.Add(key)
	}

	for key, _ := range s2 {
		unionSet.Add(key)
	}

	return unionSet
}
