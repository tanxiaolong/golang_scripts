package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "helo", "???")
	cmd.Run()
	var out bytes.Buffer
	cmd.Stdout = &out
	fmt.Println(out.String())
}
