package main

import (
	"errors"
)

func divide(x, y float64) (float64, error) {
	var err error = errors.New("no dividing by 0")
	if y == 0 {
		return 0, err
	}
	return x / y, nil
}
