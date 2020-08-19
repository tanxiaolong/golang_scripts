package main

import (
	"encoding/json"
	"fmt"
	_ "os"
	"reflect"
)

type ABC struct {
	A string `json:"a,required"`
	B int    `json:"b,required"`
}

func main() {
	at := &ABC{}
	str := "{\"a\":\"a\",\"b\":1}"
	rv2 := reflect.ValueOf(at)
	fmt.Println("rv of vt2:", rv2)
	fmt.Println("rt of at", reflect.TypeOf(at))
	fmt.Println("rt of rv2", reflect.TypeOf(rv2))
	fmt.Println("elem of rv2", rv2.Elem())
	//fmt.Println("elem of elem of rv2", rv2.Elem().Elem())
	kind2 := rv2.Kind()
	fmt.Println("kind of at", kind2)
	rv := reflect.ValueOf(&at)
	fmt.Println("rv of &vt:", rv)
	fmt.Println("rt of &at", reflect.TypeOf(&at))
	fmt.Println("rt of rv", reflect.TypeOf(rv))
	fmt.Println("elem of rv", rv.Elem())
	fmt.Println("elem of elem of rv", rv.Elem().Elem())
	kind := rv.Kind()
	fmt.Println("kind of &at:", kind)
	fmt.Println("reflect.Ptr:", reflect.Ptr)
	fmt.Println(rv.IsNil())
	err := json.Unmarshal([]byte(str), &at)
	fmt.Printf("%+v\n", at)
	fmt.Println(err)

	a := "{\"c\":123}"
	err = json.Unmarshal([]byte(a), at)
	fmt.Printf("\nerr::::::: %+v\n", err)
}
