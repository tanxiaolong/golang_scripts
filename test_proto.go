package main

import (
	"fmt"
	"google.golang.org/protobuf/runtime/protoiface"

	"git.code.oa.com/trpcprotocol/wesee_live/tag_live_daemon"
	"google.golang.org/protobuf/proto"
)

func main() {
	req := &tag_live_daemon.Test3Req{A: 16382, F: 2}

	// "github.com/golang/protobuf/proto" 可以用下面方式设置顺序
	//buf := proto.NewBuffer(nil)
	//buf.SetDeterministic(true)
	//buf.Marshal(req)
	//fmt.Printf("%v\n", buf.Bytes())

	// "google.golang.org/protobuf/proto" 是用下面的方式设置顺序
	//mo := &proto.MarshalOptions{Deterministic:true}
	//mo.Marshal(req)


	bytess, _ := proto.Marshal(req)
	fmt.Printf("%v\n\n", bytess)
	messageStruct := req.ProtoReflect()
	fmt.Printf("message: %#v\n\n", messageStruct)
	methods := messageStruct.ProtoMethods()
	fmt.Printf("methods: %#v\n\n",methods)
	size := methods.Size
	fmt.Printf("size out: %#v\n\n",size(protoiface.SizeInput{
		Message: messageStruct,
		Flags:   protoiface.MarshalInput{
			Message: messageStruct,
			Buf:     nil,
		}.Flags,
	}))
	fmt.Println()
	var b []byte
	fmt.Println(cap(b),len(b))

	noUnkeyedLiterals := methods.NoUnkeyedLiterals
	initFunc := methods.CheckInitialized
	rlt, err := initFunc(protoiface.CheckInitializedInput{
		NoUnkeyedLiterals: noUnkeyedLiterals,
		Message:           messageStruct,
	})
	fmt.Printf("%#v, err: %v\n\n", rlt, err)

	var c = make([]byte,1024)
	marshalOutput, err := methods.Marshal(protoiface.MarshalInput{
		NoUnkeyedLiterals: noUnkeyedLiterals,
		Message:           messageStruct,
		Buf:               c,
		Flags:             6,
	})
	fmt.Printf("marshalOutput: %v\n\n",marshalOutput.Buf)
}
