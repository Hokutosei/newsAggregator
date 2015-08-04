package routers

import (
	"web_apps/news_aggregator/modules/controllers"
	"web_apps/news_aggregator/modules/core/authentication"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// SetAuthenticationRoutes set authentication to the routes
func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")
	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	return router
}
