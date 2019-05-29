package main

import (
	"fmt"
	_ "os"
	"time"
)

func main() {
	start := "2019-04-19 00:00"
	end := "2019-04-25 23:00"
	showType := 2
	reverse := -1 // positive order
	reverse = 1
	timeRange, timeMap := DateRange(start, end, showType, reverse)
	fmt.Println(timeRange)
	fmt.Println(timeMap)

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
