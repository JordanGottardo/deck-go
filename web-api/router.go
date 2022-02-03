package main

import "net/http"

type Router interface {
	Post(uri string, f func(resp http.ResponseWriter, req *http.Request))
	Get(uri string, f func(resp http.ResponseWriter, req *http.Request))
	Serve(port string)
}
