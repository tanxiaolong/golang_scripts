package main

import (
	"gopkg.in/robfig/cron.v3"
	"log"
	"os"
)

func main() {
	i := 0
	c := cron.New()
	c.ErrorLog = log.New(os.Stdout, "cron: ", log.LstdFlags)
	spec := "CRON_TZ= 0 12 * * *"
	spec = "0 12 * * *"
	spec = "0 0 12 * * *"
	//spec = "* * * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	select {}
}
