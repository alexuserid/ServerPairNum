package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	first   bool
	second  = make(chan bool, 1)
	randStr string
)

func handler(w http.ResponseWriter, r *http.Request) {
	if first {
		second <- true

	} else {
		timeout := time.After(10 * time.Second)
		rand.Seed(time.Now().Unix())
		randStr = strconv.FormatInt(int64(rand.Intn(1000)), 10)
		first = true

		select {
		case <-second:
		case <-timeout:
			w.Write([]byte("Timeout. No more connected users."))
			first = false
			return
		}
	}
	w.Write([]byte(randStr))
	first = false
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
