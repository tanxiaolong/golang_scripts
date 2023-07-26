package main

import "fmt"
import "time"

func main() {
	fmt.Println("vim-go")
	nowTs := time.Now().Unix()
	now := time.Unix(nowTs, 0)
	wDay := (int(now.Weekday()) + 6) % 7
	mondy := now.AddDate(0, 0, -wDay)
	mondy = time.Date(mondy.Year(), mondy.Month(), mondy.Day(), 0, 0, 0, 0, mondy.Location())
	fmt.Println(mondy.Unix())
}
