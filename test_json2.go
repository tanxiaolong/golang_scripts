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
}
