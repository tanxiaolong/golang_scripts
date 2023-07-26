package main

import "fmt"
import "net/http"
import "io/ioutil"
import "net/url"
import "encoding/base64"

func main() {
	fmt.Println("vim-go")
	picture_url := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	rsp, err := http.Get(picture_url)
	if err != nil {
		fmt.Println(err)
	}
	image, _ := ioutil.ReadAll(rsp.Body)
	image_value, err2 := url.Parse(base64.StdEncoding.EncodeToString(image))
	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Println("base64: ", image_value.String())
}
