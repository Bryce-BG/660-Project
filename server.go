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
		for x := range in { //perminitly open pipe essentually waiting for a resposne and then appending value
			sum += x
			out <- sum
		}
	}()
}

var resultIDMx = 1

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

func (mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// var serverUrl string = "localhost"
	resultIDMx += 1
	measurementId := resultIDMx
	var imageUrl = "https://www.w3schools.com/tags/smiley.gif"
	//get task
	// get country (from IP)
	//create db object for resuilt

	temp := fmt.Sprintf(`<script id="sc1" type="text/javascript">
		var CensorshipObject = new Object();

		CensorshipObject.sendSuccess = function()
		{
		console.log("image was loaded correctly");
		// this.submitResult("success");
		}
		CensorshipObject.sendFailure = function()
		{
		console.log("image failed to load correctly");
		// this.submitResult("failure");
		}
		CensorshipObject.measure = function() {
		var img = new Image(); // width, height values are optional params
		img.src = "%s";
		img.onerror = CensorshipObject.sendFailure;
		img.onload = CensorshipObject.sendSuccess;
		// img.css('display', 'none');   //hide the image we are loading

		CensorshipObject.myID = %d;
		document.getElementById("footer").appendChild(img);

		}
		// CensorshipObject.measure();//actually call the function to execute measurement task
		</script>`, imageUrl, measurementId)

	// temp := fmt.Sprintf(`<img  src = "%s"> <script id="sc1" type="text/javascript">function showHint() {
	// 	console.log("executed this script");}</script>`, imageUrl) //###########just load basic image and test if console prints

	// serverUrl, measurementId, imageUrl)
	// temp := `<div id ='borderedDiv' ></div>` //TRY X
	// fmt.Printf(temp)
	fmt.Fprintf(w, temp)
	//
}

func main() {
	var print = fmt.Println
	print("server is running open on: localhost:8888")
	go intHandlerHelper()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./")))
	mux.Handle("/task.js", mockHandler{})
	// mux.Handle("/submit#id#", nil)
	// mux.Handle("/stats/", nil)

	http.HandleFunc("/", intHandler)
	http.ListenAndServeTLS(":8888", "server.crt",
		"server.key", mux)

}
