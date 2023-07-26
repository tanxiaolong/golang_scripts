package main

import "fmt"
import "encoding/json"
import "sync"

type TabInfo struct {
	ID         int64  `json:"id" db:"id"`
	TabName    string `json:"tab_name" db:"tab_name"`
	TabTags    string `json:"tab_tag" db:"tab_tag"`
	TagIDs     []int  `json:"tag_id"`
	PriorityID int    `json:"order_number" db:"order_number"`
	Status     int    `json:"status" db:"status"`
}

func main() {
	tabs := []*TabInfo{
		{
			ID:      1,
			TabName: "1",
			TabTags: `{"42":"生活","46":"舞蹈","61":"绝地求生","65":"单机热游"}`,
		},
		{
			ID:      2,
			TabName: "2",
			TabTags: `{"55":"生活","44":"舞蹈","33":"绝地求生","22":"单机热游"}`,
		},
		{ID: 3,
			TabName: "3",
			TabTags: `{"32":"生活","12":"舞蹈"}`,
		},
		{ID: 3,
			TabName: "4",
			TabTags: `{"76":"生活","98":"舞蹈"}`,
		},
	}
	for i := range tabs {
		fmt.Printf("before:i %+v\n", tabs[i])
	}

	parseTagIDs(&tabs)

	for i := range tabs {
		fmt.Printf("after: %+v\n", tabs[i])
	}
}

func parseTagIDs(tabs *[]*TabInfo) {
	length := len(*tabs)
	wg := sync.WaitGroup{}
	for i := 0; i < length; i++ {
		wg.Add(1)
		go func(idx int, tab *TabInfo) {
			defer wg.Done()
			tags := tab.TabTags
			tmp := map[int]string{}
			_ = json.Unmarshal([]byte(tags), &tmp)
			var tagsID []int
			for k, _ := range tmp {
				tagsID = append(tagsID, k)
			}
			tab.TagIDs = tagsID
		}(i, (*tabs)[i])
	}
	wg.Wait()
}
