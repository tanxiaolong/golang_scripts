package main

import "fmt"
import "strconv"
import "time"

func main() {
	fmt.Println("vim-go")
	p := ParseDatetime(2019101009, -1)
	fmt.Println(p)
	p = ParseDatetime("2019-04-13 12:00", -1)
	fmt.Println(p)
	p = ParseDatetime("2019-04-13", -1)
	fmt.Println(p)
}
func ParseDatetime(t interface{}, withFormat int) string {
	switch t := t.(type) {
	case int:
		tStr := strconv.Itoa(t)
		lt := len(tStr)
		if lt != 8 && lt != 10 {
			return ""
		}
		var m map[int]map[int][][]interface{}
		if lt == 8 {
			m = map[int]map[int][][]interface{}{
				8: { // day
					1:  {{"%s-%s-%s"}, {tStr[:4], tStr[4:6], tStr[6:8]}}, // with format
					-1: {{"%s-%s-%s"}, {tStr[:4], tStr[4:6], tStr[6:8]}}, // without format
				},
			}
		}
		if lt == 10 {
			m = map[int]map[int][][]interface{}{
				10: { // hour
					1:  {{"%s-%s-%s %s:00 ~ %s:59"}, {tStr[:4], tStr[4:6], tStr[6:8], tStr[8:10], tStr[8:10]}}, // with format
					-1: {{"%s-%s-%s %s:00"}, {tStr[:4], tStr[4:6], tStr[6:8], tStr[8:10]}},                     // without format
				},
			}
		}
		return fmt.Sprintf(m[lt][withFormat][0][0].(string), m[lt][withFormat][1]...)
	case string:
		if len(t) > 16 {
			return ""
		}
		_, err := time.Parse("2006-01-02 15:04", t)
		if err != nil {
			return ""
		}

		return fmt.Sprintf("%s ~ %s:59", t, t[11:13])
	}
	return ""
}
