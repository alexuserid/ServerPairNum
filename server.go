package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	cc = make(chan string)
	counter = make(chan int64)
)

func count() {
	var i int64
	for {
		i++
		counter <-i
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := <-counter

	timeout := time.After(10 * time.Second)
	randStr := strconv.FormatInt(int64(rand.Intn(1000)), 10)
	select {
	case ans := <-cc:
		w.Write([]byte(ans))
	case cc <- randStr:
		w.Write([]byte(randStr))
	case <-timeout:
		w.Write([]byte("Timeout. No more connected users."))
	}
	w.Write([]byte("\n"))
}

func main() {
	go count()
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":2080", nil))
}
