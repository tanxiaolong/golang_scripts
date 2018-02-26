//go 1.10

package main


import (
	"fmt"
	"strings"
)

func main(){
	var b strings.Builder
	b.WriteString("hello world")
	fmt.Printf("%+v\n",b)
	fmt.Printf("%+v\n",b.String())
	fmt.Println(b.Len())
	b.Grow(10)
	fmt.Println(b.Len())
	b.Reset()
	fmt.Println(b.buf)
	fmt.Println(b.String())
}
