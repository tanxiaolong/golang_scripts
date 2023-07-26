package main

import "fmt"
import "sort"

type Stu struct {
	Age    int
	Height int
	D      float64
}

type SortStu []Stu

func main() {
	students := []Stu{
		{17, 178, 1.2}, {18, 180, 3.4}, {19, 175, 5.6}, {19, 167, 7.8}, {17, 180, 9.0},
	}
	fmt.Printf("no order: %+v\n", students)

	var sortStu SortStu = students
	sort.Slice(sortStu, func(i, j int) bool {
		if sortStu[i].Age > sortStu[j].Age {
			return true
		}
		if sortStu[i].Age < sortStu[j].Age {
			return false
		}
		return sortStu[i].Height > sortStu[j].Height
	})
	fmt.Printf("order by age, then height: %+v\n", students)

	sort.Slice(sortStu, func(i, j int) bool {
		if sortStu[i].Height > sortStu[j].Height {
			return true
		}
		if sortStu[i].Height < sortStu[j].Height {
			return false
		}
		return sortStu[i].Age > sortStu[j].Age
	})
	fmt.Printf("order by height, then age: %+v\n", students)

	sort.Slice(sortStu, func(i, j int) bool {
		if sortStu[i].Height < sortStu[j].Height {
			return true
		}
		return sortStu[i].Age < sortStu[j].Age
	})
	fmt.Printf("order by height asc: %+v\n", students)

	student := []*Stu{}

	a := &Stu{17, 178, 123.4}
	b := &Stu{18, 180, 98.5}
	c := &Stu{19, 175, 912.3}
	student = append(student, a, b, c)
	fmt.Printf("before: %+v\n", student)
	sort.Slice(student, func(i, j int) bool {
		return student[i].D > student[j].D
	})
	fmt.Printf("points: %+v\n", student)

	var af float64 = 0.543508
	var bf float64 = 1030.719249
	fmt.Printf("cmp: %v\n", af < bf)
}
