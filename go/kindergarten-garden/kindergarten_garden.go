package kindergarten

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

// Define the Garden type here.
type Garden struct {
	ChildrenPlants []ChildrenPlants
}

type ChildrenPlants struct {
	ChildrenName string
	PlantNames   []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {

	// * duplicate namecheck
	if err := DuplicateNameCheck(children); err != nil {
		return nil, err
	}

	// * check the diagram format
	if err := DiagramCheck(diagram); err != nil {
		return nil, err
	}

	// * create the garden
	newChildren := make([]string, len(children))
	copy(newChildren, children)
	sort.Slice(newChildren, func(i, j int) bool {
		return newChildren[i][0] < newChildren[j][0]
	})

	garden := Garden{}
	rows := strings.Split(diagram, "\n")
	place := 0
	for _, child := range newChildren {
		plants := []string{}
		for _, row := range rows[1:] {
			if len(row) > place {
				plants = append(plants, treeMaps[string(row[place])])
				plants = append(plants, treeMaps[string(row[place+1])])
			}
		}

		chld := ChildrenPlants{
			ChildrenName: child,
			PlantNames:   plants,
		}
		garden.ChildrenPlants = append(garden.ChildrenPlants, chld)
		place += 2
	}
	return &garden, nil
}

var treeMaps = map[string]string{
	"G": "grass",
	"C": "clover",
	"R": "radishes",
	"V": "violets",
}

func (g *Garden) Plants(child string) ([]string, bool) {
	for _, inChild := range g.ChildrenPlants {
		if inChild.ChildrenName == child {
			return inChild.PlantNames, true
		}
	}
	return nil, false
}

func DuplicateNameCheck(children []string) error {

	present := make(map[string]bool)

	for _, child := range children {
		if present[child] {
			return fmt.Errorf("duplicate name found: %s", child)
		} else {
			present[child] = true
		}
	}

	return nil
}

func DiagramCheck(diagram string) error {
	// * it must start with \n
	if !strings.HasPrefix(diagram, "\n") {
		return fmt.Errorf("diagram must start with '\n'")
	}

	// Split the string by newline
	parts := strings.Split(diagram, "\n")

	// * after \n code character must be equal
	// * after \n code character cannot be odd
	// Initialize the character count
	charCount := 0
	// Count characters between newlines
	for index := 1; index < len(parts)-1; index++ {
		// if consecutive part has same length or not
		if len(parts[index]) != len(parts[index+1]) {
			return fmt.Errorf("row mismatch")
		}
		charCount += len(parts[index])
	}
	// * Check if the character count is odd
	if charCount%2 != 0 {
		return errors.New("odd number of characters between newlines")
	}

	// * allowed character code --> G,C,R,V
	if !strings.Contains(diagram, "G") && !strings.Contains(diagram, "C") && !strings.Contains(diagram, "R") && !strings.Contains(diagram, "V") {
		return errors.New("G,C,R,V must be one of present, invalid code character")
	}
	return nil
}
