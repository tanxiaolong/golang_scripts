package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	link := `https://ilive.qq.com/base/h5/pendant_container.html?activity_id=4&channel_plan_id=2&h5Url=https://isee.weishi.qq.com/iseev2/1/uwBm5DGq4/index.html?_wwv=4096`
	link = `weishi://now_live?roomid=110011384&video_url=rtmp%3A%2F%2Fws.6721.liveplay.now.qq.com%2Flive%2F6721_0208e705789756c9390821d17f917b58_hd2kfp15%3FtxSecret%3Db8e855f8ff14efc3cbcefd30ba4860e1%26txTime%3D608B1A9A`
	link = ` weishi://now_live?roomid=1361155217&source=401": first path segment in URL cannot contain colon`
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	//u.Scheme = "https"
	//u.Host = "google.com"
	q := u.Query()
	q.Set("channel_plan_id", "1")
	q.Set("channel_person_id", "iqwu8127")
	q.Set("q", "golang")
	rlt := q.Get("roomid")
	fmt.Println("rlt", rlt)
	//q.Add("activity_id","5")
	fmt.Printf("%+v\n", u.String())
	u.RawQuery = q.Encode()
	fmt.Println(u.String())

}
