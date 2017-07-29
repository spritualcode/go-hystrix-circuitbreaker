package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/spritualcode/go-hystrix-circuitbreaker/circuitbreaker"
)

func TestingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Sample to pass some custom header for the request
	var headerMap map[string]string
	headerMap = make(map[string]string)
	headerMap["Content-Type"] = "application/json"

	//PLEASE Assign Correct URL in order to see it working....
	var URL string
	URL = ""

	// If Sample API expect any post params , then please create object and replace with {"ev"} with valid struct
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode("ev")
	post_param, _ := json.Marshal("ev")

	responseStr, errchan := circuitbreaker.HystrixPostAsynch("API", URL, bytes.NewBuffer(post_param), headerMap)
	// Block until we have a result or an error. Reciever for channles sent back by Hytrix Call.
	select {
	case result := <-responseStr:
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)
		return
	case err := <-errchan:
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusOK)
		return
	}
}
