package main

import "fmt"
import "net/url"

func main() {
	webActURL := `https://abc.com/base/h5/pendant_container.html?task_h5_Url=https://asd.com/1/i6whSxRI9/index.html?_wwv=4096&needlogin=1&task_id=10`
	fmt.Println("old:", webActURL)
	webActURL = makeWebActURL(webActURL, "123", "456", 789)
	fmt.Println("new:", webActURL)
}

func makeWebActURL(webActURL string, slideID, pid string, planID int64) string {
	u, err := url.Parse(webActURL)
	if err != nil || u == nil {
		return webActURL
	}
	q := u.Query()
	q.Set("channel_room_slide_id", slideID)
	q.Set("channel_room_plan_id", fmt.Sprintf("%d", planID))
	q.Set("channel_room_person_id", pid)

	u.RawQuery = q.Encode()
	link := u.String()
	return link
}
