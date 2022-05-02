package thefarm

import "errors"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	// panic("Please implement DivideFood")
	if cows <= 0 {
		return 0, errors.New("division by zero")
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
