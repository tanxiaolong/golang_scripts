package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `txl:"hello"`
	Age string `txl:"world"`
}

func main(){
	tanxiaolong := &User{
		Name : "tanxiaolong",
		Age : "16",
	}
	fmt.Println(tanxiaolong)
	txlMap := Struct2Map(*tanxiaolong,"")
	fmt.Println(txlMap)
	txlMapTag := Struct2Map(*tanxiaolong,"txl")
	fmt.Println(txlMapTag)
	txlMapWrongTag := Struct2Map(*tanxiaolong,"txl2")
	fmt.Println(txlMapWrongTag)
}

func Struct2Map(obj interface{},tag string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if tag != "" {
			value, exists := t.Field(i).Tag.Lookup(tag)
			if exists {
				data[value] = v.Field(i).Interface()
			}
		}else {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return data
}
