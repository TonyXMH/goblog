package service

import "github.com/gorilla/mux"

func NewRouter()*mux.Router  {
	router:=mux.NewRouter().StrictSlash(true)

	for _,route:=range routes{
		router.Path(route.Pattern).Methods(route.Method).Name(route.Name).HandlerFunc(route.HandlerFunc)
	}
	return router
}
