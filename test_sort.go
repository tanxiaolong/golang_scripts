package main

import "fmt"
import "sort"

type Stu struct {
	Age    int
	Height int
}

type SortStu []Stu

func main() {
	students := []Stu{
		{17, 178}, {18, 180}, {19, 175}, {19, 167}, {17, 180},
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
}
