package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/particles", handle)
	err := http.ListenAndServe(":8080", nil)
	println(err.Error())
}

func handle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "particle-sim.html")
}
