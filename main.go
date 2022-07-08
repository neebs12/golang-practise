package main

import "fmt"

type Student struct {
	name          string
	grades        []int
	isFourthGrade bool
}

// Assign struct methods
func (s Student) getName() string {
	return s.name
}

func (s *Student) setGrades(grades []int) {
	// <--- note: is a pointer! this is to change values within the struct!
	s.grades = grades
}

func (s Student) getAverageGrade() float32 {
	var sum int
	for _, grade := range s.grades {
		sum += grade
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	std1 := Student{"Jason", []int{11, 22, 33, 44}, true}
	std2 := Student{"Jimmy", []int{88, 99, 42, 60}, false}
	fmt.Println(std1.getName(), std1.getAverageGrade(), std1)
	fmt.Println(std2.getName(), std2.getAverageGrade(), std2)

	std1.setGrades([]int{66, 77, 88})
	fmt.Println(std1.getName(), std1.getAverageGrade(), std1)
}
