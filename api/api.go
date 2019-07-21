package api

import (
	"fmt"

	"github.com/gorilla/mux"
)

func NewApi() *mux.Router {
	router := mux.NewRouter()
	handler := NumberHandler{}
	router.Path(fmt.Sprintf("/convert/{%s}", URL_NUMBER_KEY)).HandlerFunc(handler.Convert)
	return router
}
