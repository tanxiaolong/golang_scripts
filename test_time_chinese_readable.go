package main

import (
	"fmt"
	"time"
)


func main() {
	t := time.Now().AddDate(0,0,-1)
	fmt.Println(readable(t))
}

func readable(t time.Time) string {
	var timeLayout= "2006-01-02 15:04:05"
	var format= "%02d:%02d:%02d"
	now := time.Now()
	gap := now.Sub(t)
	todayHousr := now.Hour()
	gapHours := int(gap.Hours())

	if gapHours < todayHousr {
		// 今天显示时分秒
		return fmt.Sprintf(format, t.Local().Hour(), t.Local().Minute(), t.Local().Second())
	} else if gapHours > todayHousr && gapHours < 48 {
		// 昨天的显示『昨天』
		return fmt.Sprintf("昨天 "+format, t.Local().Hour(), t.Local().Minute(), t.Local().Second())
	} else if gapHours > 48 && gapHours < 72 {
		// 前天的显示『前天』
		return fmt.Sprintf("前天 "+format, t.Local().Hour(), t.Local().Minute(), t.Local().Second())
	} else {
		// 一周以内，前天以后的显示日期
		return t.Format(timeLayout)
	}
}
