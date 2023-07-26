package main

import "net/url"
import "fmt"

func main() {
	//webActURL := `https://ilive.qq.com/base/h5/pendant_container.html?task_h5_Url=https://test-isee.weishi.qq.com/iseev2/1633/VmIIxlUG6/index.html?_wwv=4096&needlogin=1&task_id=20&test=test`
	webActURL := `http:abc.com?pid=78530&roseId=40573484&targetid=6954293470&source=9&tabid=0&tabtype=1`
	u, err := url.Parse(webActURL)
	str := u.String()
	fmt.Println(str, err)
	urlParam := u.RawQuery
	fmt.Println("urlParam:", urlParam)
	m, err := url.ParseQuery(urlParam)
	if err == nil {
		fmt.Printf("\n\n map: %v\n\n", m)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
