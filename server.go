package main

import (
	"io"
	"net/http"
	"strconv"
)

var in = make(chan int)
var out = make(chan int)
var sum int

func intHandlerHelper() {
	go func() {
		for x := range in {
			sum += x
			out <- sum
		}
	}()
}

func intHandler(w http.ResponseWriter, r *http.Request) {
	// USAGE: access http://127.0.0.1:8888/?x=[number]
	xString := r.URL.Query()["x"][0]
	x, err := strconv.Atoi(xString)
	if err == nil {
		in <- x
		result := <-out
		io.WriteString(w, strconv.Itoa(result))
	} else {
		io.WriteString(w, "INVALID")
	}

}

func main() {
	go intHandlerHelper()
	http.HandleFunc("/", intHandler)
	http.ListenAndServeTLS(":8888", "server.crt",
		"server.key", nil)
}
