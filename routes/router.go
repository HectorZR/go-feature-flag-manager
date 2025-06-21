package routes

import (
	"fmt"
	"net/http"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)

func formatRouteString(route string, method HttpMethod) string {
	return fmt.Sprintf("%s %s", method, route)
}

type Router struct{}

func (r Router) Get(route string, handler http.HandlerFunc) {
	http.HandleFunc(formatRouteString(route, GET), handler)
}

func (r Router) Post(route string, handler http.HandlerFunc) {
	http.HandleFunc(formatRouteString(route, POST), handler)
}

func (r Router) Put(route string, handler http.HandlerFunc) {
	http.HandleFunc(formatRouteString(route, PUT), handler)
}

func (r Router) Patch(route string, handler http.HandlerFunc) {
	http.HandleFunc(formatRouteString(route, PATCH), handler)
}

func (r Router) Delete(route string, handler http.HandlerFunc) {
	http.HandleFunc(formatRouteString(route, DELETE), handler)
}
