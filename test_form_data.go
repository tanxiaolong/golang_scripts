package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "mime/multipart"
    "net/http"
)

func main() {
    postData := make(map[string]map[string]string)
    postData["media"] = map[string]string{
		"contentType":"image/png",
		"value":"",
	}
    PostWithFormData("POST",url ,&postData)
}


func PostWithFormData(method, url string, postData *map[string]map[string]string){
    body := new(bytes.Buffer)
    w := multipart.NewWriter(body)
    for k,v :=  range *postData{
        w.WriteField(k, v)
    }
    w.Close()
    req, _ := http.NewRequest(method, url, body)
    req.Header.Set("Content-Type", w.FormDataContentType())
    resp, _ := http.DefaultClient.Do(req)
    data, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    fmt.Println(resp.StatusCode)
    fmt.Printf("%s", data)
}
