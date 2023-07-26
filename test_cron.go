package main

import (
	"github.com/robfig/cron"
	"log"
	"os"
)

func main() {
	i := 0
	c := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	//c := cron.New()
	spec := "CRON_TZ= 0 12 * * *"
	spec = "0 12 * * *"
	//spec = "* * * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	select {}
}
