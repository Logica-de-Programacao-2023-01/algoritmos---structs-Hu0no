package main

import "fmt"

type Worker struct {
	Name string
	Wage float64
	Age  uint
}

func IncreaseWage(worker *Worker, percent int) float64 {
	increase := worker.Wage * float64(percent) / 100
	worker.Wage += increase
	return worker.Wage
}
func DecreaseWage(worker *Worker, percent int) float64 {
	decrease := worker.Wage * float64(percent) / 100
	worker.Wage -= decrease
	return worker.Wage
}
func CalculateTimeOfService(worker *Worker) uint {
	time := worker.Age - 18
	return time
}
func main() {
	worker := Worker{
		Name: "Marcola",
		Wage: 5000,
		Age:  25,
	}
	percent := 50
	fmt.Print(DecreaseWage(&worker, percent))
}
