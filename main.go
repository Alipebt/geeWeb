package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)
	log.Fatal(http.ListenAndServe(":9999",nil))
}

// handler echoes r.URL.Path

	func indexHandler(w http.ResponseWriter,req *http.Request){
		fmt.Fprintf(w,"hello")
	}

