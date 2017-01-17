package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Reponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

// Docs
// https://golang.org/pkg/net/http
// https://golang.org/pkg/io/#Writer

// This is our function we are going to use to handle the request
// All handlers need to accept two arguments
// 1. Is the ResponseWriter interface, we use this to write a reponse back to the client
// 2. Is the Reponse struct which holds useful information about the request headers, method, url etc
func hello(w http.ResponseWriter, r *http.Request) {
	// We use the standard libaries WriteStirng function to write back to the ResponseWriter interface
	// See docs above

	name := r.FormValue("name")
	var ResponseCode int
	var ResponseStr string

	switch nameLength := len(name); {
	case nameLength == 0:
		ResponseCode = 403
		ResponseStr = "Please set a name"
	case nameLength < 2:
		ResponseCode = 403
		ResponseStr = "Please supply a name longer than 1 character"
	default:
		ResponseCode = 200
		ResponseStr = fmt.Sprintf("%s %s", "hello", r.FormValue("name"))
	}

	respond(w, ResponseStr, ResponseCode)
}

func respond(w http.ResponseWriter, ResponseStr string, ResponseCode int) {
	res := Reponse{
		Code:   ResponseCode,
		Result: ResponseStr,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	// Add ads the function thats going to handle that response
	http.HandleFunc("/", hello)
	// Starts the web server
	// The first argument in this method is the port you want your server to run on
	// The second is a handler. However we have already added this in the line above so we just pass in nil
	servicePort := flag.Int("service_port", 8000, "Port that the webservice will run on")
	flag.Parse()

	http.ListenAndServe(":"+strconv.Itoa(*servicePort), nil)
}
