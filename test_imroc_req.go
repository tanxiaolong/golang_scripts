package main

import (
	"fmt"
	"github.com/imroc/req"
	"log"
)

func main() {
	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic YWRtaW46YWRtaW4=",
	}
	param := req.Param{
		"name": "imroc",
		"cmd":  "add",
	}
	// 只有url必选，其它参数都是可选
	r, err := req.Post("http://foo.bar/api", header, req.BodyJSON(param))
	if err != nil {
		log.Fatal(err)
	}
	r.ToJSON(&foo)       // 响应体转成对象
	log.Printf("%+v", r) // 打印详细信息
}
