package main


import "fmt"

func main(){

	i := 3
	var h, l uint8 = uint8(i>>8), uint8(i&0xff)
	fmt.Println(i,h,l)

}
