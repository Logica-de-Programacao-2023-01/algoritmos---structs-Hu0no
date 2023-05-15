package main

type Triangulo struct {
	basis  float64
	height float64
}

func area(triangulo Triangulo) float64 {
	area := (triangulo.height * triangulo.basis) / 2
	return area
}
