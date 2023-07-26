package main

import "fmt"
import "encoding/json"
import "reflect"

func main() {

	str := `{"sid":["123-5"],"keys":["baidumedia_1643811406708927859"],"metas":{"baidumedia_1643811406708927859":{"is_top":1}}}`
	test := &RequestGetCountByKeyBatch{}
	json.Unmarshal([]byte(str), &test)
	fmt.Printf("%+v", test)

	fmt.Println()

	test2 := &RequestGetCountByKeyBatch2{}
	json.Unmarshal([]byte(str), &test2)
	fmt.Printf("%+v", test2)

	fmt.Println()
	fmt.Println(reflect.TypeOf(test2.Metas))

	fmt.Println("test2.Metas[\"baidumedia_1643811406708927859\"]:", test2.Metas["baidumedia_1643811406708927859"])

	fmt.Println("test2.Metas[\"baidumedia_1643811406708927859\"][\"is_top\"]:", test2.Metas["baidumedia_1643811406708927859"]["is_top"])
}

type RequestGetCountByKeyBatch struct {
	Sid   []string    `json:"sid"`
	Keys  []string    `json:"keys"`
	Metas interface{} `json:"metas"`
}

type RequestGetCountByKeyBatch2 struct {
	Sid   []string                          `json:"sid"`
	Keys  []string                          `json:"keys"`
	Metas map[string]map[string]interface{} `json:"metas"`
}
