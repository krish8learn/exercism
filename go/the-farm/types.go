package thefarm

import (
	"errors"
)

// This file contains types used in the exercise but should not be modified.

// WeightFodder returns the amount of available fodder.
type WeightFodder interface {
	FodderAmount() (float64, error)
}

// ErrScaleMalfunction indicates an error with the scale.
var ErrScaleMalfunction = errors.New("sensor error")

type farm struct {
	fodder float64
	err    error
}

func (f farm) FodderAmount() (float64, error) {
	if f.fodder > 0 && f.err == ErrScaleMalfunction{
		return f.fodder, ErrScaleMalfunction
	}
	return 0, f.err
}
