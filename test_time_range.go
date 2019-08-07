package main

import (
	"fmt"
	_ "os"
	"time"
)

func main() {
	//2019-07-18 00:00 2019-07-18 23:00 2
	start := "2019-07-17 00:00"
	end := "2019-07-18 23:00"
	showType := 2
	reverse := "asc"
	isTimeReverse := 1
	timeRange, _ := DateRange(start, end, showType, reverse, isTimeReverse)
	fmt.Println(timeRange)
}

func DateRange(dateStart, dateEnd string, showType int, reverse string, isTimeReverse int) ([]string, map[string]int) {
	var dateMap = map[int]map[string]map[string]interface{}{
		1: { // 分时
			"desc": { // true reverse
				"offset":       "-1h",
				"position":     16,
				"format":       "2006-01-02 15:00",
				"dateStart":    dateEnd,
				"dateEnd":      dateStart,
				"position_asc": 10,
				"offset_asc":   "1h",
			},
			"asc": { // po
				"offset":       "1h",
				"position":     16,
				"format":       "2006-01-02 15:00",
				"dateStart":    dateStart,
				"dateEnd":      dateEnd,
				"position_asc": 10,
				"offset_asc":   "-1h",
			},
		},
		2: { // 分日
			"desc": { // true reverse
				"offset":    "-24h",
				"position":  10,
				"format":    "2006-01-02",
				"dateStart": dateEnd,
				"dateEnd":   dateStart,
			},
			"asc": { // false reverse not
				"offset":    "24h",
				"position":  10,
				"format":    "2006-01-02",
				"dateStart": dateStart,
				"dateEnd":   dateEnd,
			},
		},
	}
	if showType == 1 && isTimeReverse == -1 {
		dateStart = dateMap[showType][reverse]["dateStart"].(string)[:dateMap[showType][reverse]["position_asc"].(int)] + " 00:00"

		dateEnd = dateMap[showType][reverse]["dateEnd"].(string)[:dateMap[showType][reverse]["position_asc"].(int)] + " 23:00"
	} else {
		dateStart = dateMap[showType][reverse]["dateStart"].(string)[:dateMap[showType][reverse]["position"].(int)]

		dateEnd = dateMap[showType][reverse]["dateEnd"].(string)[:dateMap[showType][reverse]["position"].(int)]
	}

	fmt.Println(dateStart, dateEnd)
	startTime, _ := time.Parse(dateMap[showType][reverse]["format"].(string), dateStart)
	startTimeStr := startTime.Format(dateMap[showType][reverse]["format"].(string))
	endTimeStr := dateEnd

	timeRange := []string{startTimeStr}
	timeMap := map[string]int{
		startTimeStr: 0,
	}
	if startTimeStr == endTimeStr {
		return timeRange, timeMap
	}
	i := 1
	for {
		// 主要考虑到分时时候，如果要求时间
		if showType == 1 && isTimeReverse == -1 && i%24 == 0 {
			tmp, _ := time.ParseDuration("-48h")
			startTime = startTime.Add(tmp)
		}
		startTimeStr = startTime.Format(dateMap[showType][reverse]["format"].(string))
		fmt.Println(1, startTimeStr)
		var t time.Duration
		// 主要考虑到分时时候，如果要求时间
		if showType == 1 && isTimeReverse == -1 {
			t, _ = time.ParseDuration(dateMap[showType][reverse]["offset_asc"].(string))
		} else {
			t, _ = time.ParseDuration(dateMap[showType][reverse]["offset"].(string))
		}
		fmt.Println(t)
		startTime = startTime.Add(t)
		startTimeStr = startTime.Format(dateMap[showType][reverse]["format"].(string))
		timeRange = append(timeRange, startTimeStr)
		timeMap[startTimeStr] = i
		i++
		if startTimeStr == endTimeStr {
			break
		}
	}
	return timeRange, timeMap
}
