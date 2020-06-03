package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Omikuji おみくじjson
type Omikuji struct {
	Omikuji string `json:"omikuji"`
}

var omikuzi = map[int]string{0: "大吉", 1: "中吉", 2: "末吉", 3: "凶"}

func handler(w http.ResponseWriter, r *http.Request) {
	item := rand.Intn(len(omikuzi))
	result, _ := omikuzi[item]
	omikujijson := &Omikuji{Omikuji: result}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(omikujijson); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, buf.String())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
