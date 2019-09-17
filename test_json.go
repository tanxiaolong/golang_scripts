package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type A struct {
	A string
}

func main() {

	test := `{\"6005856\":[\"TG101\",\"TG102\"],\"6005909\":[\"TG101\",\"TG102\",\"TG103\",\"TG104\",\"TG101\",\"TG102\",\"TG101\",\"TG102\",\"TG101\",\"TG102\",\"TG101\",\"TG102\",\"TG101\",\"TG102\",\"TG101\",\"TG102\",\"TG101\",\"TG102\",\"TG101\",\"TG102\"]}`
	fmt.Println(test)
	test = strings.ReplaceAll(test, "\\", "")
	fmt.Println(test)
	testArr := map[string][]string{}
	errr := json.Unmarshal([]byte(test), &testArr)
	fmt.Println(testArr, errr)
	os.Exit(200)

	//b := &A{}
	//a := ""
	//err := json.Unmarshal([]byte(a), &b)
	//fmt.Println(err)

	at := &AdReq{}
	err := json.Unmarshal([]byte(str), &at)
	fmt.Printf("%+v\n", at)
	fmt.Println(err)
}

var str = `{"name":"白色","campaign_id":273,"budget_limit_type":1,"budget":100000,"budget_money":100,"bid":10000,"bid_money":10,"ad_target":{"age":[],"genders":[],"geo_locations":{},"isp":[],"device_brand":[],"user_os":[],"user_network":[],"user_interest":[],"register_time_range":[],"filter_installed_app":100001,"white_list":[1,2,34]},"ad_schedule":{"start_date":"","end_date":"","hours_week":{},"schedule_type":1},"id":380,"customer_id":306205,"conf_status":2,"need_delete":0}`

var strr = `{"white_list":[1,2]}`

type AdReq struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CustomerID int64  `json:"customer_id"`
	CampaignID int64  `json:"campaign_id"`

	BudgetLimitType int   `json:"budget_limit_type"`
	Budget          int64 `json:"budget"`
	Bid             int64 `json:"bid"`
	ConfStatus      int   `json:"conf_status"`

	AdTarget   AdTargetReq   `json:"ad_target"`
	AdSchedule AdScheduleReq `json:"ad_schedule"`

	NeedDelete int `json:"need_delete"`
}

type AdScheduleReq struct {
	StartDate string           `json:"start_date"`
	EndDate   string           `json:"end_date"`
	HoursWeek map[string][]int `json:"hours_week"`

	ScheduleType int `json:"schedule_type"`
}

type AdTargetReq struct {
	Age                []int            `json:"age"`
	Genders            []int            `json:"genders"`
	GeoLocations       map[string][]int `json:"geo_locations"`
	Isp                []int            `json:"isp"`
	DeviceBrand        []int            `json:"device_brand"`
	UserOs             []int            `json:"user_os"`
	UserNetwork        []int            `json:"user_network"`
	UserInterest       []int            `json:"user_interest"`
	RegisterTimeRange  []int            `json:"register_time_range"`
	FilterInstalledApp int              `json:"filter_installed_app"`
	WriteList          []int            `json:"white_list"`
}
type b struct {
}
