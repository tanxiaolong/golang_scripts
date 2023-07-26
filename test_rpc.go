package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

type Watcher int

func (w *Watcher) GetInfo(arg int, result *int) error {
	*result = 1
	return nil
}

func main() {

	http.HandleFunc("/rpc", RpcTest)

	watcher := new(Watcher)
	rpc.Register(watcher)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败，端口可能已经被占用")
	}
	fmt.Println("正在监听1234端口")
	http.Serve(l, nil)
}

func RpcTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>rpc test</body></html>")
}
