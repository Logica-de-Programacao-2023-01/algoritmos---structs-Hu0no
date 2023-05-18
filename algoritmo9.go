package main

import "fmt"

type Student struct {
	Name   string
	Age    uint
	Grades []float64
}

func AddGrades(student Student, x float64) []float64 {
	student.Grades = append(student.Grades, x)
	return student.Grades
}
func RemoveGrades(student Student, x float64) []float64 {
	newgrades := []float64{}
	for _, grade := range student.Grades {
		if grade != x {
			newgrades = append(newgrades, grade)
		}
	}
	student.Grades = newgrades
	return student.Grades
}
func Average(student Student) float64 {
	soma := 0.0
	for _, grade := range student.Grades {
		soma += grade
	}
	average := soma / float64(len(student.Grades))
	fmt.Println(average)
	return average
}
func PrintInfo(student Student) int {
	fmt.Println("Student name: ", student.Name)
	fmt.Println("Student Age: ", student.Age)
	fmt.Println("Student Average: ", Average(student))
	return 0
}
