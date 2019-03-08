package main

import (
	"fmt"
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
	// USAGE: access https://127.0.0.1:8888/?x=[number]
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

type mockHandler struct {
}

func (m *mockHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	go intHandlerHelper()

	mux := http.NewServeMux()
	// mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(staticPath))))
	ts := NewTasksServer()
	mux.Handle("/task.js", ts)
	mux.Handle("/submit", nil)
	mux.Handle("/stats/", nil)

	http.HandleFunc("/", intHandler)
	http.ListenAndServeTLS(":8888", "server.crt",
		"server.key", mux)

}
