package main

import (
	"net/http"
)

type Route struct {
	Pattern       string
	GETHandler    func(w http.ResponseWriter, r *http.Request)
	POSTHandler   func(w http.ResponseWriter, r *http.Request)
	PUTHandler    func(w http.ResponseWriter, r *http.Request)
	DELETEHandler func(w http.ResponseWriter, r *http.Request)
}

type Routes []Route

func (rs Routes) findOrNew(pattern string) Route {
	for _, r := range rs {
		if pattern == r.Pattern {
			return r
		}
	}
	return Route{Pattern: pattern}
}

func (rs Routes) insert(newRoute Route) {
	for i, r := range rs {
		if newRoute.Pattern == r.Pattern {
			rs[i] = newRoute
			return
		}
	}
	routes = append(routes, newRoute)
}

var routes Routes

//DefRoute creates a Route and adds it to our routes array
func DefRoute(method string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {

	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		panic("Invalid HTTP method registering")
	}

	route := routes.findOrNew(pattern)
	switch method {
	case "GET":
		route.GETHandler = handler
	case "POST":
		route.POSTHandler = handler
	case "PUT":
		route.PUTHandler = handler
	case "DELETE":
		route.DELETEHandler = handler
	}
	routes.insert(route)
}

//RouteAll translates all the routes created with DefRoute to go's API. Calls GC after registering it
func RouteAll() {
	for _, route := range routes {
		http.HandleFunc(route.Pattern, func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				if route.GETHandler != nil {
					route.GETHandler(w, r)
				} else {
					http.NotFound(w, r)
				}
			case "POST":
				if route.POSTHandler != nil {
					route.POSTHandler(w, r)
				} else {
					http.NotFound(w, r)
				}
			case "PUT":
				if route.PUTHandler != nil {
					route.PUTHandler(w, r)
				} else {
					http.NotFound(w, r)
				}
			case "DELETE":
				if route.DELETEHandler != nil {
					route.DELETEHandler(w, r)
				} else {
					http.NotFound(w, r)
				}
			}
		})
	}
	//garbage collect
	routes = nil
}
