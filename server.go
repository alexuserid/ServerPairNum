package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type t struct {
	str string
	n int64
	ch chan int64
}

var (
	cc = make(chan t)
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
	inst := t{str: strconv.FormatInt(int64(rand.Intn(1000)), 10), n: <-counter}
	timeout := time.After(10 * time.Second)

	select {
	case ans := <-cc:
		w.Write([]byte(ans.str))
	case cc <- inst:
		w.Write([]byte(inst.str))
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
