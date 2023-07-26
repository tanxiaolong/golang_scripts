package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"git.code.oa.com/trpcprotocol/wesee_live/tag_live_daemon"
	"google.golang.org/protobuf/proto"
	"math"
)

func main() {
	req := &tag_live_daemon.Test3Req{
		A: 1,
		B: 1.23,
		C: 16382,
		D: map[string]string{"a": "b", "c": "d", "e": "f"},
		E: "proto",
		F: []string{"l", "m", "n"},
		G: 2,
		H: &tag_live_daemon.Nested{
			Aa: 3,
		},
	}

	fixedStr := "abc"
	fmt.Println("字符串：", fixedStr)
	fixedBytes := md5.Sum([]byte(fixedStr))
	fmt.Println("固定字符串，一次：", fixedBytes)
	fixedBytes = md5.Sum([]byte(fixedStr))
	fmt.Println("固定字符串，二次：", fixedBytes)

	bytes, _ := proto.Marshal(req)
	dynamicBytes := md5.Sum(bytes)
	fmt.Println("动态字符串，一次：", dynamicBytes)
	bytes, _ = proto.Marshal(req)
	dynamicBytes = md5.Sum(bytes)
	fmt.Println("动态字符串，二次：", dynamicBytes)
	mo := &proto.MarshalOptions{Deterministic: true}
	rlt, _ := mo.Marshal(req)
	fmt.Println("stable", rlt)
	rlt, _ = mo.Marshal(req)
	fmt.Println("stable2", rlt)
	req = &tag_live_daemon.Test3Req{
		A: math.MaxInt32,
		G: 2,
		C: 16382,
		B: 1.23,
		E: "proto",
		F: []string{"l", "m", "n"},
		H: &tag_live_daemon.Nested{
			Aa: 5,
			Bb: "abc",
		},
	}
	fmt.Println(math.MaxInt32)
	bytes, _ = proto.Marshal(req)
	fmt.Println("pb字节数组: ", bytes)
	bytes, _ = json.Marshal(req)
	fmt.Println("json字节数组：", bytes)
}
