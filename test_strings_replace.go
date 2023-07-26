package main

import "strings"
import "fmt"

func main(){

str := "今天|是个|好日||子"
str = strings.ReplaceAll(str,"|","\\|")
fmt.Println(str)
}
