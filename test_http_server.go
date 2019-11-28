package main

import (
	"encoding/json"
	"fmt"
	//"html"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(502)
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		rsp,_ := json.Marshal(map[string]int{
			"err_code":200,
		})
		fmt.Fprintf(w, "%s", string(rsp))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}