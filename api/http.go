package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	romanNumbers "github.com/chonla/roman-number-go"
	"github.com/gorilla/mux"
)

const (
	URL_NUMBER_KEY = "number"
)

var nonascii = map[string]int{ // (JF) currently useless
	"ↁ": 5000,
	"ↂ": 10000,
}

var errorMap = map[int]string{
	100001: "Could not serialize response",
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Trace   string `json:"trace"`
}

type ApiResponse struct {
	Roman  string `json:"roman,omitempty"`
	Arabic int    `json:"arabic,omitempty"`
}

type NumberHandler struct {
}

func (h *NumberHandler) Convert(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numberAsStr, ok := vars[URL_NUMBER_KEY]
	if !ok {
		http.Error(rw, "No number specified", 400)
		return
	}
	roman := romanNumbers.NewRoman()
	parsedNumber, err := strconv.Atoi(numberAsStr)
	resp := ApiResponse{}
	if err == nil {
		resp.Roman = roman.ToRoman(parsedNumber)
	} else {
		resp.Arabic = roman.ToNumber(numberAsStr)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		// (JF) Could also us log.Fatalf or other logging library
		log.Printf("error during serialization of response (input: %q): %s", numberAsStr, err)
		bErr, err := json.Marshal(ApiError{
			Code:    100001,
			Message: errorMap[100001],
			Trace:   "TODO USE OPENTRACING AND JAEGER",
		})
		if bErr != nil {
			panic(err)
			return
		}
		rw.WriteHeader(500)
		rw.Write(bErr) // (JF) ignore bytes written in this example
		return
	}
	rw.Write(b) // (JF) ignore bytes written in this example
}
