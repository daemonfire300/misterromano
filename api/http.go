package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	100002: "Invalid input",
}

var validRomans = "MDCLXVI"

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

func validRoman(in string) bool {
	return strings.ContainsAny(in, validRomans)
}

func writeErr(rw http.ResponseWriter, errCode int) {
	bErr, err := json.Marshal(ApiError{
		Code:    errCode,
		Message: errorMap[errCode],
		Trace:   "TODO USE OPENTRACING AND JAEGER",
	})
	if err != nil {
		panic(err)
		return
	}
	rw.WriteHeader(500)
	rw.Write(bErr) // (JF) ignore bytes written in this example
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
		if !validRoman(numberAsStr) {
			log.Printf("error not a valid input (input: %q)", numberAsStr)
			writeErr(rw, 100002)
			return
		}
		resp.Arabic = roman.ToNumber(numberAsStr)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		// (JF) Could also us log.Fatalf or other logging library
		log.Printf("error during serialization of response (input: %q): %s", numberAsStr, err)
		writeErr(rw, 100001)
		return
	}
	rw.Write(b) // (JF) ignore bytes written in this example
}
