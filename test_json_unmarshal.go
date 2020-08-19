package main

import "fmt"
import _ "runtime/debug"
import "encoding/json"
import "reflect"

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

	str = `{"msg_id":159430674800001314,"seq_id":1594306738580,"msg_type":11,"content":"null","type":11,"sender":1142002,"receiver":1161876,"all_unread_count":2,"version_id":1594306748002,"without_sender":0,"without_recipient":0}`
	data := SendMsgPush{}
	err = json.Unmarshal([]byte(str), &data)
	fmt.Printf("unmarshal data:%v\n", err)
	fmt.Printf("data:%+v\n", data)

	tm := &TextMsg{}
	fmt.Println("data.Content's type:", reflect.TypeOf(data.Content))
	txt := data.Content.(string)
	fmt.Println("xxx:", data.Content)
	err = json.Unmarshal([]byte(txt), &tm)
	fmt.Printf("unmarshal tm: %v\n", err)
	fmt.Printf("tm:%+v\n", tm)

}

type SendMsgPush struct {
	MsgID            uint64      `json:"msg_id"`
	SeqID            int64       `json:"seq_id"`
	MsgType          int8        `json:"msg_type"`
	Content          interface{} `json:"content"`
	Type             int8        `json:"type"`
	SenderUID        uint64      `json:"sender"`
	PeerUID          uint64      `json:"peer_id"`
	ReceiveUID       uint64      `json:"receiver"`
	AllUnreadCount   int64       `json:"all_unread_count"`
	VersionID        int64       `json:"version_id"`
	UpdateTime       int64       `json:"update_time"`
	WithoutSender    int8        `json:"without_sender"`
	WithoutRecipient int8        `json:"without_recipient"`
	Extra            string      `json:"extra"`
}

type TextMsg struct {
	TextContent struct {
		Content string `json:"content"`
	} `json:"text_content"`
}
