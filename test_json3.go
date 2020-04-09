package main

import "fmt"
import "encoding/json"

func main() {
	fmt.Println("vim-go")
	str := `{"location":{"business":"","department":"nvwa","start":"","end":"","url":"/api/v1/imcenter/message/*","desc":"中台私聊汇聚层发消息接口","check_session":0,"http_method":"","visitor_limit":0,"check_account_safe":0,"limit_ip":"","tokens":[],"default_data":"","status":1,"url_id":"","check_signature":0,"app_name":"","upstream":[{"app":"nvwa","department":"nvwa","service_name":"imcenter.pm.buz","url_id":"","white_list":[],"uid_suffix_range":{"start":"","end":""}}]}`
	rlt := map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &rlt)
	fmt.Println(err)

}
