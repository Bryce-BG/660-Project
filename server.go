package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	MeasurementId int    `json:"measurement_id"`
	Result        string `json:"result_status"`
}

var in = make(chan int)
var out = make(chan int)
var sum int
var imgTemplate string
var db *Database
var err error
var online bool

func intHandlerHelper() {
	go func() {
		for x := range in { //perminitly open pipe essentually waiting for a resposne and then appending value
			sum += x
			out <- sum
		}
	}()
}

var resultIDMx = 1

//UNUSED
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

type submitHandler struct {
}

func (mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.UserAgent(), r.Host, r.RemoteAddr)
	// var serverUrl string = "localhost"

	// resultIDMx += 1
	// measurementId := resultIDMx
	// var imageUrl = "https://www.w3schools.com/tags/smiley.gif"
	taskID, imageURL, _ := db.OfferRandomTask()
	// measurementID, _ := db.AddResultEntry(taskID, "8.8.8.8", "mars", "moon", time.Time{}, "ghost", 0.0)
	// TODO : placeholder for id.
	country, city, latitude, longtitude := getLocaleData("8.8.8.8")
	measurementID, err := db.AddResultEntry(taskID, r.RemoteAddr, r.UserAgent(), country, city, latitude, longtitude)
	if err != nil {
		log.Fatal(err)
	}

	//get task
	// get country (from IP)
	//create db object for resuilt

	temp := fmt.Sprintf(imgTemplate, imageURL, measurementID)
	fmt.Fprintf(w, temp)

}

func (submitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		// fmt.Println(r.UserAgent)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		t := Response{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			log.Fatal(err)
		}
		db.UpdateResult(t.MeasurementId, t.Result)
		// fmt.Println("result: " + t.Result + " measurementid: " + strconv.Itoa(t.MeasurementId))

	}
}

func main() {
	b, _ := ioutil.ReadFile("task_templates/img.js")
	imgTemplate = string(b)

	db, err = Initialize()
	if err != nil {
		log.Println(err)
	}

	port, isSet := os.LookupEnv("PORT")
	online = isSet
	if isSet == false {
		port = "8888"
	}

	var print = fmt.Println
	print("server is running open on: localhost:" + port)
	go intHandlerHelper()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./")))
	mux.Handle("/task.js", mockHandler{})
	mux.Handle("/submit", submitHandler{})
	// mux.Handle("/stats/", nil)

	http.HandleFunc("/", intHandler)
	// http.ListenAndServe(":"+port, mux)
	http.ListenAndServeTLS(":"+port, "server.crt", "server.key", mux)

}
