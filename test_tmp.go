package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("2016-08-03", "2016-08-20")
	a, b := DateRange("2016-08-03", "2016-08-20", 2)
	fmt.Println(a)
	fmt.Println(b)
	cd := newChartData(a)
	fmt.Println(cd)
	gg := map[int]string{
		1: "2",
		3: "4",
	}
	fmt.Println(gg[1])
	fmt.Println(gg[5])
}

func DateRange(dateStart, dateEnd string, showType int) ([]string, map[string]int) {

	var timeRange []string
	var timeMap map[string]int
	if strings.Contains(dateStart, " ") {
		dateStart = dateStart[:10]
		dateEnd = dateEnd[:10]
	}
	startDay, _ := time.Parse("2006-01-02", dateStart)
	fmt.Println("startDay:", startDay)
	switch showType {
	case 1:
		{ // 分时
			// todo
		}
	case 2:
		{ // 分日

			timeRange = append(timeRange, dateStart)
			timeMap = map[string]int{
				dateStart: 1,
			}
			for {
				startDay = startDay.AddDate(0, 0, 1)
				nextDay := startDay.Format("2006-01-02")
				if nextDay == dateEnd {
					break
				}
				timeRange = append(timeRange, nextDay)
				timeMap[nextDay] = 1
			}
			timeRange = append(timeRange, dateEnd)
		}
	}
	return timeRange, timeMap
}
func newChartData(time []string) *ChartData {
	l := len(time)
	cd := &ChartData{
		PV:    make([]int, l),
		Click: make([]int, l),
		CTR:   make([]float32, l),
		ACPE:  make([]float32, l),
		ECPM:  make([]float32, l),
		Cost:  make([]float32, l),
		Time:  time,
	}
	return cd
}

type ChartData struct {
	PV    []int     `json:"pv"`
	Click []int     `json:"click"`
	CTR   []float32 `json:"ctr"`
	ACPE  []float32 `json:"acpe"`
	ECPM  []float32 `json:"ecpm"`
	Cost  []float32 `json:"cost"`
	Time  []string  `json:"time"`
}
