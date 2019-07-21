package main

import (
	"log"
	"net/http"
	"os"

	"github.com/daemonfire300/misterromano/api"
)

func main() {
	log.SetOutput(os.Stdout)
	router := api.NewApi()
	http.ListenAndServe(":8080", router)
}
