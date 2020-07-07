package main

import "fmt"
import "encoding/json"

type HobbyInfo struct {
	Type      string   `json:"type"`
	Icon      string   `json:"icon"`
	Color     string   `json:"color"`
	TextColor string   `json:"text_color"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
}

// 多人响应体
type MUserProfileRsp struct {
	DMErr  int        `json:"dm_error"`
	ErrMsg string     `json:"error_msg"`
	Data   []UserInfo `json:"data"`
}

type UserInfo struct {
	Uid         uint64                   `json:"uid"`
	Nick        string                   `json:"nick"`
	Portrait    string                   `json:"portrait"`
	Gender      int8                     `json:"gender"`
	Birth       string                   `json:"birth"`
	Age         int32                    `json:"age"`
	Signature   string                   `json:"signature"`
	Media       []map[string]interface{} `json:"media,omitempty"`
	Back        map[string]interface{}   `json:"back,omitempty"`
	Hobby       []HobbyInfo              `json:"hobby,omitempty"`
	Anonymous   int8                     `json:"anonymous"`
	Description string                   `json:"description"`
	Profession  string                   `json:"profession"`
	HomeTown    string                   `json:"hometown"`
	ID          int64                    `json:"id"`
	VIPInfo     struct {
		VIPType   int `json:"vip_type"`
		VIPStatus int `json:"vip_status"`
	} `json:"vip_info"`
}

func main() {
	str := `{"dm_error":0,"error_msg":"请求成功","data":[{"description":"947322","hometown":"947322","id":947322,"portrait":"947322","profession":"947322"},{"description":"794409","hometown":"794409","id":794409,"portrait":"794409","profession":"794409"},{"description":"605181","hometown":"605181","id":605181,"portrait":"605181"}]}`
	dst := &MUserProfileRsp{}
	str = ""
	err := json.Unmarshal([]byte(str), &dst)
	fmt.Println(err)
	fmt.Println("vim-go", dst)
}
