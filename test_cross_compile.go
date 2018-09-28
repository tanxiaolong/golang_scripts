package main

import "fmt"

/*

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test_cross_compile.go// 4 win

GOARCH    target architecture 
GOOS      target operate system

OS                   ARCH                          OS version

linux                386 / amd64 / arm             >= Linux 2.6

darwin               386 / amd64                   OS X (Snow Leopard + Lion)

freebsd              386 / amd64                   >= FreeBSD 7

windows              386 / amd64                   >= Windows 2000
*/

func main(){
	fmt.Println("hello Windows! -- from centos");
}
