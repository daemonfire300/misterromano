package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dataRoman = map[string]int{
	"MCMLXXXIV": 1984,
	"I":         1,
	"V":         5,
	"IV":        4,
	"VI":        6,
	"X":         10,
}

func TestNumberHandler_ConvertToArabic(t *testing.T) {
	srv := NewApi()
	for input, expected := range dataRoman {
		t.Run(fmt.Sprintf("Roman Number %s is correctly converted to %d", input, expected), func(t *testing.T) {
			rq := httptest.NewRequest("GET", "http://example.org/convert/"+input, nil)
			rspRecorder := httptest.NewRecorder()
			srv.ServeHTTP(rspRecorder, rq)
			resp := rspRecorder.Result() // (JF) usually as a real client you would call defer resp.Body.Close() in order to avoid open connections
			// i.e., resource leak
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("error during test: %s", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("expected http resonse code <%d> but got <%d>", http.StatusOK, resp.StatusCode)
			}
			apiResp := ApiResponse{}
			err = json.Unmarshal(content, &apiResp)
			if err != nil {
				t.Fatalf("error during test: %s", err)
			}
			if apiResp.Arabic != expected {
				t.Fatalf("expected result <%d> but got <%d>", expected, apiResp.Arabic)
			}
		})
	}

}

func TestNumberHandler_ConvertToRoman(t *testing.T) {
	srv := NewApi()
	for expected, input := range dataRoman {
		t.Run(fmt.Sprintf("Arabic Number %d is correctly converted to %s", input, expected), func(t *testing.T) {
			rq := httptest.NewRequest("GET", fmt.Sprintf("http://example.org/convert/%d", input), nil)
			rspRecorder := httptest.NewRecorder()
			srv.ServeHTTP(rspRecorder, rq)
			resp := rspRecorder.Result() // (JF) usually as a real client you would call defer resp.Body.Close() in order to avoid open connections
			// i.e., resource leak
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("error during test: %s", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("expected http resonse code <%d> but got <%d>", http.StatusOK, resp.StatusCode)
			}
			apiResp := ApiResponse{}
			err = json.Unmarshal(content, &apiResp)
			if err != nil {
				t.Fatalf("error during test: %s", err)
			}
			if apiResp.Roman != expected {
				t.Fatalf("expected result <%s> but got <%s>", expected, apiResp.Roman)
			}
		})
	}

}
