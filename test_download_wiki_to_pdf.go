package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	flag.Parse()
	pageID := flag.Arg(0)
	log.Println("page id:",pageID)
	//pageID := "80216974"
	url := "https://wiki.inkept.cn/spaces/flyingpdf/pdfpageexport.action?pageId="+pageID
	req(pageID,url)
	fmt.Println("vim-go")
}

func req(pageID,url string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		// 治疗302等跳转
		CheckRedirect: checkRedirect,
		// 治疗x509: certificate signed by unknown authority
		Transport: tr,
		Timeout:   10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`)
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", `_ga=GA1.2.1185115270.1595338255; user=%257B%2522id%2522%253A100739%252C%2522identification%2522%253A%2522dc%253Dinke%2522%252C%2522loginName%2522%253A%2522tanxiaolong%2522%252C%2522password%2522%253A%2522%2522%252C%2522userType%2522%253A%25220%2522%252C%2522delFlag%2522%253A%25220%2522%252C%2522realName%2522%253A%2522%25E8%25B0%25AD%25E6%2599%2593%25E9%25BE%2599%2522%252C%2522description%2522%253A%2522%2522%252C%2522department%2522%253A%2522%25E4%25B8%25AD%25E5%258F%25B0%25E7%25A0%2594%25E5%258F%2591%25E9%2583%25A8%2522%252C%2522phone%2522%253A%2522%2522%252C%2522mail%2522%253A%2522tanxiaolong%2540inke.cn%2522%252C%2522employeeNumber%2522%253A%2522%2522%252C%2522timestamp%2522%253Anull%252C%2522jobType%2522%253A%2522%25E6%2596%25B0%25E4%25BA%25A7%25E5%2593%2581%25E4%25B8%25AD%25E5%25BF%2583%2522%252C%2522authKey%2522%253A%2522%2522%252C%2522roles%2522%253A%255B%255D%252C%2522userRoles%2522%253A%255B%255D%252C%2522roleNamesStr%2522%253A%2522%25E5%2588%259D%25E5%25A7%258B%25E8%25A7%2592%25E8%2589%25B2%2522%257D; token=STXxFKHzYXXDUvNOzxMnbdZugkuzPdhpRvd; TOKEN=STXJMRhTNYlsLcSdlQYMdyiEwzxuHksMvBO; JSESSIONID=0BD182F69157A2ED2AD596F14D3703E3; _gid=GA1.2.1483048351.1599551901`)
	req.Header.Add("Host", "wiki.inkept.cn")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Referer", "https://wiki.inkept.cn/pages/viewpage.action?pageId="+pageID)
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(resp, err)
		return
	}
	defer resp.Body.Close()

	attachment := resp.Header.Get("Content-Disposition")
	log.Println("attachment",attachment)
	log.Println("attachment:::",reflect.TypeOf(attachment))


	body, err := ioutil.ReadAll(resp.Body)
	fileName := strings.Split(strings.Split(attachment, ";")[1],"=")[1]

	writeToFile(fileName[1:len(fileName)-1], string(body))
}


func checkRedirect(req *http.Request,via []*http.Request)error{
	return nil
}

func writeToFile(filename, msg string)  {
	location := "/Users/tanxiaolong/Downloads/wiki/"+filename
	log.Println("location:",location)
	f, err := os.OpenFile(location, os.O_WRONLY|os.O_CREATE, 0644)
	_, err = f.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
	}
}