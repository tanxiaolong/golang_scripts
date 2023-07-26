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
	spec = "@every 1s"      // execute at every second
	spec = "1 * * * *"      // execute at every first second of m
	spec = "1/1 * * * *"    // like @every 1s
	spec = "0/10 * * * * *" // every 10s execute once, starts from 0, not the time you execute
	//spec = "* * * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	select {}
}
