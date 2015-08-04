package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes init routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
