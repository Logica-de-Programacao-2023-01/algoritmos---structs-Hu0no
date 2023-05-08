package main

import (
	"math"
)

type Circulo struct {
	radius float64
}

func area(circle Circulo) float64 {
	return math.Pi * (circle.radius * circle.radius)
}
