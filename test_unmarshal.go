package main

import "encoding/json"
import "fmt"
import "reflect"

func main() {
	jsonStr := []byte(`{"age":1}`)
	var value map[string]interface{}
	json.Unmarshal(jsonStr, &value)
	age := value["age"]
	fmt.Println(reflect.TypeOf(value))
	fmt.Println(reflect.TypeOf(age))
	//float64
}
