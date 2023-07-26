package main

import "fmt"
import "net/url"

func main() {
	str := buildDataKey("tabid=100&tabtype=111&a=1&bb=2")
	fmt.Println(str)
}

func buildDataKey(origin string) string {
	// 加个假前缀，好解析
	fakePrefix := "http://abc.com?"
	u, err := url.Parse(fakePrefix + origin)
	if err != nil || u == nil {
		return origin
	}
	q := u.Query()
	q.Del("tabid")
	q.Del("tabtype")
	q.Set("tabtype", "1")
	q.Set("tabid", "0")

	str := q.Encode()

	return str
}
