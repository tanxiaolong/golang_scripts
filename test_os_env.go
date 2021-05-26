package main

import "fmt"
import "os"

func main() {
	fmt.Println(os.Getenv("TRPC_LOG_TRACE"))
}
