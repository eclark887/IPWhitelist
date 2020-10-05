package main

import (
	"strings"
	"net/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
		grpcServer.ServeHTTP(w, r)
		} else {
		httpHandler.ServeHTTP(w, r)
		}
	})
}


func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		homePage,
	},
	Route{
		"Healthz",
		"GET",
		"/healthz",
		healthz,
	},
	Route{
		"Healthz",
		"GET",
		"/livez",
		healthz,
	},
	Route{
		"Healthz",
		"GET",
		"/readyz",
		healthz,
	},
	Route{
		"IPWhitelist",
		"POST",
		"/api/v1/ipwhitelist",
		postIPWhitelist,
	},
}
