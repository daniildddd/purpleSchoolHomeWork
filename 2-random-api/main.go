package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

// домашка по курсу осталось лишь протестить и сделать merge(перед этим на ветку вторую закинуть)
func randHandler(w http.ResponseWriter, r *http.Request) {
	numb := (rand.Intn(6) + 1)
	w.Write([]byte(strconv.Itoa(numb)))
}

func main() {
	http.HandleFunc("/", randHandler)
	http.ListenAndServe(":9091", nil)
}
