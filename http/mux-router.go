package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	router = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (r *muxRouter) GET(uri string, fn func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(uri, fn).Methods("GET")
}

func (r *muxRouter) POST(uri string, fn func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(uri, fn).Methods("POST")
}

func (r *muxRouter) SERVE(port string) error {
	log.Println("Server running on port: ", port)
	err := http.ListenAndServe(port, router)
	return err
}
