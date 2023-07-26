package main

import "flag"
import "fmt"

// 定义命令行参数对应的变量，这三个变量都是指针类型
var cliName = flag.String("name", "nick", "Input Your Name")
var cliAge = flag.Int("age", 28, "Input Your Age")
var cliGender = flag.String("gender", "male", "Input Your Gender")

// 定义一个值类型的命令行参数变量，在 Init() 函数中对其初始化
// 因此，命令行参数对应变量的定义和初始化是可以分开的
var cliFlag int
func Init() {
    flag.IntVar(&cliFlag, "flagname", 1234, "Just for demo")
}

func main(){
  fmt.Println(cliName,cliAge,cliGender,cliFlag)
}
