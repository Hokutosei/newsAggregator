package routers

import (
	"web_apps/news_aggregator/modules/controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
