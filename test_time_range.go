package main

import (
	"fmt"
	_ "os"
	"time"
)

func main() {
	start := "2019-04-19 00:00"
	end := "2019-04-25 23:00"
	showType := 1
	reverse := -1 // positive order
	reverse = 1
	isTimeReverse := -1
	timeRange, _ := DateRange(start, end, showType, reverse)
	fmt.Println(timeRange)
	//fmt.Println(timeMap)
	fmt.Println("=======================")
	timeRange, _ = DateRange2(start, end, showType, reverse, isTimeReverse)
	fmt.Println(timeRange)
	//fmt.Println(timeMap)
	fmt.Println("++++++++++++++++++++++++")
	reverseStr := "asc"
	isTimeReverse = 1
	fmt.Println(start, end)
	timeRange, _ = DateRange3(start, end, showType, reverseStr, isTimeReverse)
	fmt.Println(timeRange)
}

func DateRange3(dateStart, dateEnd string, showType int, reverse string, isTimeReverse int) ([]string, map[string]int) {
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
				"test":         "-48h",
			},
			"asc": { // po
				"offset":       "1h",
				"position":     16,
				"format":       "2006-01-02 15:00",
				"dateStart":    dateStart,
				"dateEnd":      dateEnd,
				"position_asc": 10,
				"offset_asc":   "-1h",
				"test":         "0h",
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

	startTime, _ := time.Parse(dateMap[showType][reverse]["format"].(string), dateStart)
	startTimeStr := startTime.Format(dateMap[showType][reverse]["format"].(string))
	endTimeStr := dateEnd

	timeRange := []string{startTimeStr}
	timeMap := map[string]int{
		startTimeStr: 0,
	}
	i := 1
	for {
		// 主要考虑到分时时候，如果要求时间
		if showType == 1 && isTimeReverse == -1 && i%24 == 0 {
			tmp, _ := time.ParseDuration("-48h")
			startTime = startTime.Add(tmp)
		}
		startTimeStr = startTime.Format(dateMap[showType][reverse]["format"].(string))
		//if startTimeStr == endTimeStr {
		//	return timeRange, timeMap
		//}
		var t time.Duration
		// 主要考虑到分时时候，如果要求时间
		if showType == 1 && isTimeReverse == -1 {
			t, _ = time.ParseDuration(dateMap[showType][reverse]["offset_asc"].(string))
		} else {
			t, _ = time.ParseDuration(dateMap[showType][reverse]["offset"].(string))
		}
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

func DateRange2(dateStart, dateEnd string, showType int, reverse int, isTimeReverse int) ([]string, map[string]int) {
	var dateMap = map[int]map[int]map[string]interface{}{
		1: { // 分时
			1: { // true reverse
				"offset":       "-1h",
				"position":     16,
				"format":       "2006-01-02 15:00",
				"dateStart":    dateEnd,
				"dateEnd":      dateStart,
				"position_asc": 10,
				"offset_asc":   "1h",
			},
			-1: { // po
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
			1: { // true reverse
				"offset":    "-24h",
				"position":  10,
				"format":    "2006-01-02",
				"dateStart": dateEnd,
				"dateEnd":   dateStart,
			},
			-1: { // false reverse not
				"offset":    "24h",
				"position":  10,
				"format":    "2006-01-02",
				"dateStart": dateStart,
				"dateEnd":   dateEnd,
			},
		},
	}
	fmt.Println("date", dateStart, dateEnd)
	if showType == 1 && isTimeReverse == -1 {
		dateStart = dateMap[showType][reverse]["dateStart"].(string)[:dateMap[showType][reverse]["position_asc"].(int)] + " 00:00"

		dateEnd = dateMap[showType][reverse]["dateEnd"].(string)[:dateMap[showType][reverse]["position_asc"].(int)] + " 23:00"
	} else {
		dateStart = dateMap[showType][reverse]["dateStart"].(string)[:dateMap[showType][reverse]["position"].(int)]

		dateEnd = dateMap[showType][reverse]["dateEnd"].(string)[:dateMap[showType][reverse]["position"].(int)]
	}
	fmt.Println("dateStart", dateStart, "dateEnd", dateEnd)

	startTime, _ := time.Parse(dateMap[showType][reverse]["format"].(string), dateStart)

	endTimeStr := dateEnd
	startTimeStr := startTime.Format(dateMap[showType][reverse]["format"].(string))

	timeRange := []string{startTimeStr}
	timeMap := map[string]int{
		startTimeStr: 0,
	}
	i := 1
	for {
		if showType == 1 && isTimeReverse == -1 && i%24 == 0 {
			fmt.Println("777", startTime)
			fmt.Println("time range ", timeRange)
			tmp, _ := time.ParseDuration("-48h")
			startTime = startTime.Add(tmp)
			//fmt.Println("6.7", startTime)
			//tmp, _ = time.ParseDuration("-24h")
			//startTime = startTime.Add(tmp)
			//fmt.Println("666", startTime)
			//os.Exit(1)
		}
		startTimeStr = startTime.Format(dateMap[showType][reverse]["format"].(string))
		if startTimeStr == endTimeStr {
			return timeRange, timeMap
		}
		var t time.Duration
		if showType == 1 && isTimeReverse == -1 {
			t, _ = time.ParseDuration(dateMap[showType][reverse]["offset_asc"].(string))
		} else {
			t, _ = time.ParseDuration(dateMap[showType][reverse]["offset"].(string))
		}
		//os.Exit(1)
		startTime = startTime.Add(t)
		nextTimeStr := startTime.Format(dateMap[showType][reverse]["format"].(string))
		timeRange = append(timeRange, nextTimeStr)
		timeMap[nextTimeStr] = i
		i++
		if nextTimeStr == endTimeStr {
			break
		}
	}
	return timeRange, timeMap
}

func DateRange(dateStart, dateEnd string, showType int, reverse int) ([]string, map[string]int) {
	var dateMap = map[int]map[int]map[string]interface{}{
		1: map[int]map[string]interface{}{
			1: map[string]interface{}{ // true reverse
				"offset":    "-1h",
				"position":  16,
				"format":    "2006-01-02 15:00",
				"dateStart": dateEnd,
				"dateEnd":   dateStart,
			},
			-1: map[string]interface{}{ // po
				"offset":    "1h",
				"position":  16,
				"format":    "2006-01-02 15:00",
				"dateStart": dateStart,
				"dateEnd":   dateEnd,
			},
		},
		2: map[int]map[string]interface{}{
			1: map[string]interface{}{ // true reverse
				"offset":    "-24h",
				"position":  10,
				"format":    "2006-01-02",
				"dateStart": dateEnd,
				"dateEnd":   dateStart,
			},
			-1: map[string]interface{}{ // false reverse not
				"offset":    "24h",
				"position":  10,
				"format":    "2006-01-02",
				"dateStart": dateStart,
				"dateEnd":   dateEnd,
			},
		},
	}

	//dateStart = dateStart[:dateMap[showType][reverse]["position"].(int)]
	dateStart = dateMap[showType][reverse]["dateStart"].(string)[:dateMap[showType][reverse]["position"].(int)]

	//dateEnd = dateEnd[:dateMap[showType][reverse]["position"].(int)]
	dateEnd = dateMap[showType][reverse]["dateEnd"].(string)[:dateMap[showType][reverse]["position"].(int)]
	fmt.Println("dateStart,dateEnd:", dateStart, dateEnd)
	startTime, _ := time.Parse(dateMap[showType][reverse]["format"].(string), dateStart)
	endTimeStr := dateEnd
	startTimeStr := startTime.Format(dateMap[showType][reverse]["format"].(string))
	timeRange := []string{startTimeStr}
	timeMap := map[string]int{
		startTimeStr: 0,
	}
	i := 1
	for {
		startTimeStr = startTime.Format(dateMap[showType][reverse]["format"].(string))
		if startTimeStr == endTimeStr {
			timeRange = append(timeRange, startTimeStr)
			timeMap[startTimeStr] = i
			return timeRange, timeMap
		}
		t, _ := time.ParseDuration(dateMap[showType][reverse]["offset"].(string))
		startTime = startTime.Add(t)
		nextTimeStr := startTime.Format(dateMap[showType][reverse]["format"].(string))
		timeRange = append(timeRange, nextTimeStr)
		timeMap[nextTimeStr] = i
		i++
		if nextTimeStr == endTimeStr {
			break
		}
	}
	return timeRange, timeMap
}
