package main
import (
	log "unknwon.dev/clog/v2"
)

func init() {
	err := log.NewConsole("你好")
	if err != nil {
		panic("unable to create new logger: " + err.Error())
	}
}

func main() {
	log.Trace("Hello %s!", "World") // YYYY/MM/DD 12:34:56 [TRACE] Hello World!
	log.Info("Hello %s!", "World")  // YYYY/MM/DD 12:34:56 [ INFO] Hello World!
	log.Warn("Hello %s!", "World")  // YYYY/MM/DD 12:34:56 [ WARN] Hello World!

	// Graceful stopping all loggers before exiting the program.
	log.Stop()
}
