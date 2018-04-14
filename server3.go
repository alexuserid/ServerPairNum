package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	cli int
)

func handler(w http.ResponseWriter, r *http.Request) {
	cli++
	fmt.Println(cli)
	time.Sleep(10 * time.Second)
	w.Write([]byte("success"))
	cli = 0
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
