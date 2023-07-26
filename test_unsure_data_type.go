package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

/*
sometimes we don't know what a field's specific value, the value could be 1 or "1",
so how we can define our strongly typed struct ?

json.RawMessage is a good choice.

*/
type MsgCollect struct {
	From json.RawMessage `json:"from"`
}

func main() {

	str := `{"from":"1"}`
	msgCol := &MsgCollect{}
	json.Unmarshal([]byte(str), &msgCol)

	fmt.Printf("%+v\n", msgCol)

	fmt.Println(string(msgCol.From))
	fmt.Println("1")

	fmt.Println("type: ", reflect.TypeOf(msgCol.From))

	// convert to int
	fromString := string(msgCol.From)
	l := len(fromString)
	fromInt, _ := strconv.Atoi(fromString[1 : l-1])
	fmt.Println(fromInt, " , type: ", reflect.TypeOf(fromInt))

	// convert to string
	fromStr := string(msgCol.From)
	fmt.Println(fromStr, " , type: ", reflect.TypeOf(fromStr))

	//======================================================
	str = `{"from":1}`
	//msgCol = &MsgCollect{}
	json.Unmarshal([]byte(str), &msgCol)

	fmt.Printf("%+v\n", msgCol)

	fmt.Println(string(msgCol.From))

	fmt.Println("type: ", reflect.TypeOf(msgCol.From))

	// convert to int
	fromInt, _ = strconv.Atoi(string(msgCol.From))
	fmt.Println(fromInt, " , type: ", reflect.TypeOf(fromInt))

	// convert to string
	fromStr = string(msgCol.From)
	fmt.Println(fromStr, " , type: ", reflect.TypeOf(fromStr))
}
