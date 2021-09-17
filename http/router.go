package router

import (
	"net/http"
)

type Router interface {
	GET(uri string, fn func(http.ResponseWriter, *http.Request))
	POST(uri string, fn func(http.ResponseWriter, *http.Request))
	SERVE(port string) error
}
