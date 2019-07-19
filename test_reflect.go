package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{
		Name: "tanxiaolong",
		Age:  30,
	}
	p = &Person{}
	t := reflect.TypeOf(p)
	fmt.Println(t)
	valOf := reflect.ValueOf(p)
	fmt.Println("reflect value of:", valOf)
	kValOf := valOf.Kind()
	fmt.Println("kind of value of:", kValOf)
	val := reflect.ValueOf(p).Elem()
	fmt.Println("reflect value of item:", val)
	structType := val.Type().Kind()
	fmt.Println("type of the struct is:", structType)
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("json"))
	}

	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("json"))
}
