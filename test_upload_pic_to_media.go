package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// you need a "Basic Auth" plugin in your wordpress
// you can get it from github.com/WP-API/Basic-Auth
// usage: go run test_upload_pic_to_media.go snoopy.jpg

func main() {
	flag.Parse()
	// pass the pic path
	path := flag.Arg(0)
	// upload it
	upload(path)
}

func upload(path string) {
 	// read the pic
	data := readPic(path)
	if data != "" {
		// post it
		httpDoPost(data)
	}else{
		fmt.Println("pic is null")
		return 
	}
}

func readPic(path string) string{
	// io.Reader
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read error")
		return ""
	}
	return string(data)
}

func httpDoPost(data string) {

	client := &http.Client{}

	url := "https://domain/wp-json/wp/v2/media/"
	//url := "https://domain/wp-json"
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	//req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	username := "abcde"
	password := "12345"
	req.SetBasicAuth(username,password)
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("content-disposition", "attachment; filename=naiwenwen.jpg")
	req.Header.Set("cache-control","no-cache")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	fmt.Println(string(body))
}


