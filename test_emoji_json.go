package main

import (
	"encoding/json"
	"fmt"
)

type Names struct {
	Name string `json:"name"`
}

func main() {

	emoji := map[string]string{"name": "😏百变大咖秀😂"}
	emojiBytes, err := json.Marshal(emoji)
	fmt.Println("json:", string(emojiBytes), err)

	n := &Names{}
	json.Unmarshal(emojiBytes, n)
	fmt.Println(n)
}
