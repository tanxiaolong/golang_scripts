package main

import (
	"net/http"
	"io"
)
func HelloHandler(w http.ResponseWriter, r *http.Request){
	str := "Hello world"
	io.WriteString(w,str)
}

func main(){
	ht := http.HandlerFunc(HelloHandler)
	if ht !=nil {
		http.Handle("/hello",ht)
	}
	http.ListenAndServe(":8090",nil)
}
