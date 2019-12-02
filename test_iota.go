package main

import "fmt"

// 消息类型
const (
	_                  = iota
	MsgTypeText        // 1文本
	MsgTypeAudio       // 2音频
	MsgTypeImage       // 3图片
	MsgTypeParent      // 4礼物
	MsgTypeRedirectUrl // 5跳转地址

	SystemMsg
)

// 消息类型
const (
	_ = iota
	MsgTypeText1
	MsgTypeAudio2
	MsgTypeImage3
	MsgTypeParent4
	MsgTypeRedirectUrl5
)

func main() {
	fmt.Println(SystemMsg, MsgTypeImage, MsgTypeImage3)
}
