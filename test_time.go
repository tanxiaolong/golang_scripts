package main

import (
	"fmt"
	"time"
)

var (
	GoDay           = "2006-01-02"
	GoBirthday      = "2006-01-02 00:00"
	GoBirthDatetime = "2006-01-02 15:04:05"
)

func main() {
	s := time.Now().Local().String()
	fmt.Println(s)
	s = time.Now().Local().Format(GoBirthDatetime)
	fmt.Println(s)
	s = time.Now().Local().Format(GoBirthday)
	fmt.Println(s)
	s = time.Now().Format(GoDay)
	fmt.Println("today date:", s)
	todayTime, _ := time.Parse("2006-01-02", s)
	fmt.Println("aaaaaaa:", todayTime.Unix()-8*3600)
	fmt.Println("aaaaaaa:", todayTime.Unix()+16*3600)
	s = time.Now().Local().AddDate(0, 0, -1).Format(GoBirthday)
	fmt.Println(s)
	fmt.Println("================")
	s = time.Now().Local().AddDate(0, 0, -1).Format(GoBirthDatetime)
	fmt.Println(s)
	s = time.Now().Local().AddDate(0, -1, 0).Format(GoBirthDatetime)
	fmt.Println(s)
	now := time.Now().Local()
	s = now.AddDate(0, 0, -now.Day()+1).Format(GoBirthDatetime)
	fmt.Println(s)
	fmt.Println("-----------------------")
	queryDateStart := now.AddDate(0, 0, -14).Format(GoBirthday)
	queryDateEnd := now.Format(GoBirthday)
	fmt.Println(queryDateStart, queryDateEnd)
	ti := "2019032421"
	fmt.Println(ti[:4])
	fmt.Println(ti[4:6])
	fmt.Println(ti[6:8])
	fmt.Println(ti[8:10])

	s = time.Now().Local().Format(GoBirthday)[:11] + "24:00"
	fmt.Println(s)

	x := "2019-04-08 14:45:56"
	p, err := time.Parse("2006-01-02", x)
	fmt.Println("123:", err)
	fmt.Println(p.Date())
	fmt.Println(x)
	fmt.Println(x[:4] + x[5:7] + x[8:10] + x[11:13] + x[14:16] + x[17:19])
	fmt.Println("55555%:", x[:10])
	rlt := p.AddDate(0, 0, 1).Format(GoDay)
	fmt.Println("rlt:", rlt)
	x = "2019-04-30"
	p, _ = time.Parse("2006-01-02", x)
	fmt.Println("p:", p)
	tmp := p.AddDate(0, 0, 1).Format(GoDay)
	fmt.Println("tmp:", tmp)
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println("now:", now.Format("2006-01-02 15:00"))
	fmt.Println("1 hour after:", hh1.Format("2006-01-02 15:00"))
	day, _ := time.ParseDuration("-24h")
	dd1 := now.Add(day)
	fmt.Println("24 hours after:", dd1.Format("2006-01-02"))
	fmt.Println("2018-09-02 13:00" < "2018-09-02 14:00")
	str := "2018-02-12 01:00"
	fmt.Println(str[:len(str)], len(str))
	fmt.Println(time.Parse("2006-01-02 15:00", str))

	t := "2019-04-14 15:00"
	fmt.Println(len(t))
	t = "2019-04-14"
	p, err1 := time.Parse("2006-01-02", t)
	fmt.Println(p)
	fmt.Println(err1)
	p, err2 := time.Parse("2006-01-02 15:04", t)
	fmt.Println("tanxiaolong", p, p.Format("2006-01-02 15:04"))
	fmt.Println(err2)
	fmt.Println("now:", p.Format("2006010215"))

	fmt.Println("========================")
	dateStart := "2019-06-13 23:00"
	_, err = time.Parse("2006-01-02 23:00", dateStart)
	fmt.Println("err", err)

	txxxxxxxxx := "2019-09-03 23:59"
	txxxxxxTime, err := time.Parse("2006-01-02 15:04", txxxxxxxxx)
	fmt.Println("txxxxxxxx:", txxxxxxTime, err)
	tSSSStr := txxxxxxTime.Format("2006010215")
	fmt.Println(tSSSStr)

	yesterday := time.Now().AddDate(0, 0, -1).Format("20060102")
	fmt.Println("yesterday:", yesterday)

	fmt.Println(time.Now().Format("20060102"))

	tt := time.Now()
	tm1 := time.Date(tt.Year(), tt.Month(), tt.Day(), 23, 59, 59, 0, tt.Location())
	fmt.Println("today's last second: ", tm1)
	fmt.Println("today's last second: ", tm1.Unix())

	xxxxxx := "2020-01-04 23:43:35"
	pp, _ := time.Parse("2006-01-02 15:04:05", xxxxxx)
	fmt.Printf("ttttttt: %v\n", pp)
	fmt.Println(pp.Unix())

	ppo := time.Unix(1578470764, 0)
	fmt.Println(ppo.Format("2006-01-02 15:04"))

}
