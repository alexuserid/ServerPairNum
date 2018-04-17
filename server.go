package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	ch              = make(chan bool, 1)
	first           bool
	second          = make(chan bool, 1)
	randStr, strRec string
	toStr           string = "Timeout. No more connected users.\n"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ch <- true
	select {
	case <-ch:
		if first {
			strRec = randStr + "\n"
			second <- true

		} else {
			timeout := time.After(10 * time.Second)
			randStr = strconv.FormatInt(int64(rand.Intn(1000)), 10)
			first = true

			select {
			case <-second:
			case <-timeout:
				strRec = toStr
			}
		}
		w.Write([]byte(strRec))
		first = false
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
