package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	slice := [][]string{{"A", "B"}}
	slice0 := []string{"3", "4"}
	slice = append(slice, slice0)
	slice1 := []string{"5", "6"}
	slice = append(slice, slice1) //slice的输出为[[A B] [3 4] [5 6]]
	w.WriteAll(slice)             //WriteAll保存slice二维数据
	w.Flush()
	fmt.Println(slice)
}
