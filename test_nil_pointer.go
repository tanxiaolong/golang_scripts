package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	//var p1 interface{}
	//p1 = nil
	//fmt.Println("4", p1.(float64))

	pp := map[string]interface{}{"1": 1, "2": "2", "3": 1.1}
	fmt.Println(pp)
	fmt.Println("before marshal", reflect.TypeOf(pp["1"]), reflect.TypeOf(pp["3"]))
	ppBytes, _ := json.Marshal(pp)
	ppp := map[string]interface{}{}
	json.Unmarshal(ppBytes, &ppp)
	fmt.Println("after marshal and unmarshal", reflect.TypeOf(ppp["1"]), reflect.TypeOf(ppp["3"]))
	pppp := map[string]int{}
	json.Unmarshal(ppBytes, &pppp)
	fmt.Println("after marshal and unmarshal", pppp)

	ppppp := map[interface{}]interface{}{}
	_, err := json.Marshal(ppppp)
	fmt.Println("map[interface{}]interface{} type can not be json marshaled", "err:", err)

	pppppp := map[interface{}]string{"1": "1", 2: "2"}
	fmt.Println("pppppp-'1'", pppppp["1"])
	fmt.Println("pppppp-'2'", pppppp["2"])
	fmt.Println("pppppp-2", pppppp[2])

	_, err = json.Marshal(pppppp)
	fmt.Println("map[interface{}]interface{} type can not be json marshaled---2", "err:", err)

	fmt.Println(pp["4"])

	var aa int64 = 11
	a := map[interface{}]interface{}{1: 111, 2: "2"}
	b := map[int64]int64{3: 3}
	if val, exists := a[float64(aa)]; !exists {
		fmt.Println("666", val)
	}
	fmt.Println("1", a[float64(aa)])
	fmt.Println("2", a[float64(aa)].(float64)) // must be checked before a convertion!!!!
	fmt.Println("3", b[int64(a[float64(aa)].(float64))])

}
