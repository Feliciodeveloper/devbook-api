package routes

import (
	"api/src/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

//Route represents all API routes
type Route struct {
	URI string
	Method string
	Function func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}
func Configure(r *mux.Router) *mux.Router{
	routes := routeUsers
	routes = append(routes,routeLogin)
	routes = append(routes,routePosts...)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI,
				middleware.Logger(middleware.Authenticate(route.Function))).
				Methods(route.Method)
		}else{
			r.HandleFunc(route.URI,middleware.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}