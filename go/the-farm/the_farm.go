package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	//  panic("Please implement DivideFood")
	if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0.0, &SillyNephewError{cows: cows}
	}
	amount, err := weightFodder.FodderAmount()
	if err != nil {
		if err != ErrScaleMalfunction {
			return 0, err
		} else if err == ErrScaleMalfunction && amount < 0 {
			return 0, errors.New("negative fodder")
		} else if err == ErrScaleMalfunction && amount > 0 {
			return (amount * 2) / float64(cows), nil
		}
	} else if amount < 0 {
		return 0, errors.New("negative fodder")
	}
	return amount / float64(cows), nil
}
