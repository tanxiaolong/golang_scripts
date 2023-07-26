package main

import (
	"fmt"
	"math"
	"reflect"
)
type buffer []byte
func main(){
	a:=3.111111111111111111111111111118
	fmt.Printf("%f\n",a)
	fmt.Printf("%.30f\n",math.Pi)
	fmt.Printf("%+v\n",math.Pi)
	fmt.Println(reflect.TypeOf(math.Pi))
	//math.Pi
	//3.14159265358979323846264338327950288419716939937510582097494459
	//3.141592653589793115997963468544
	//3.141592653589793
}