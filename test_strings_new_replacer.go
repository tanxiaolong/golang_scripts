package main

import "fmt"
import "strings"

func main() {
	stream_id_format := "s-{keyinfo}-{liveid}"
	//stream_id_format = "zegotest-{appid}-s-{keyinfo}-{liveid}"

	srp := strings.NewReplacer("{appid}", "gmu", "{keyinfo}", "666", "{liveid}", "1234567")
	var streamID []string
	streamID = append(streamID, srp.Replace(stream_id_format))
	fmt.Println(streamID)
	fmt.Println(streamID[0])
}
