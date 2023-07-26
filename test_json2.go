package main

import "fmt"
import "encoding/json"

func main() {

	str := `{"user_info": "{"uid": 1297965, "nick": "\u738b\u5c0f\u6263", "sex": 1, "portrait": "http://img.ikstatic.cn/MTU1ODM0MjI2ODA3NSMzNzIjanBn.jpg"}"}`
	//str = `{"user_info":"{"uid":1297965,"nick":"\u738b\u5c0f\u6263","sex":1,"portrait":"http://img.ikstatic.cn/MTU1ODM0MjI2ODA3NSMzNzIjanBn.jpg"}"}`
	//str = `{"user_info":"{\"uid\"\:1297965,\"nick"\:\"牛B\",\"sex\"\:1,\"portrait\"\:\"http://img.ikstatic.cn/MTU1ODM0MjI2ODA3NSMzNzIjanBn.jpg\"}"}`
	str = `{"uid": 1297965, "nick": "\u738b\u5c0f\u6263", "sex": 1, "portrait": "http://img.ikstatic.cn/MTU1ODM0MjI2ODA3NSMzNzIjanBn.jpg"}`
	str = `{"user_info": "{\"uid\": 1297965, \"nick\": \"\u738b\u5c0f\u6263\", \"sex\": 1, \"portrait\": \"http://img.ikstatic.cn/MTU1ODM0MjI2ODA3NSMzNzIjanBn.jpg\"}"}`
	rlt := map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &rlt)
	fmt.Printf("%+v\n", err)
	fmt.Println(rlt)

	test := &Test{
		A: 1,
		B: str,
	}

	testBytes, _ := json.Marshal(test)
	fmt.Println(string(testBytes))

	test2 := &Test{}
	json.Unmarshal(testBytes, test2)
	asd := map[string]interface{}{}
	json.Unmarshal([]byte(test2.B), &asd)
	fmt.Println(asd)
}

type Test struct {
	A int    `json:"a"`
	B string `json:"b"`
}
