package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var omikuzi = map[int]string{0: "大吉", 1: "中吉", 2: "末吉", 3: "凶"}

func handler(w http.ResponseWriter, r *http.Request) {
	item := rand.Intn(len(omikuzi))
	result, _ := omikuzi[item]
	fmt.Fprint(w, result)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
