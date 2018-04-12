package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const	wait time.Duration = 10

var (
	clients int = 0
	rndmS string
)

func handler(w http.ResponseWriter, r *http.Request) {
	clients++
	fmt.Println(clients)

	time.Sleep(wait * time.Second)

	if clients == 2 {
		rndm := rand.Intn(1000)
		rndmS = strconv.FormatInt(int64(rndm), 10)
		http.Redirect(w, r, "/ok/", http.StatusSeeOther)
	} else {
		w.Write([]byte("Only two clients can be served at once"))
		clients = 0
	}
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(rndmS))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ok/", okHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
