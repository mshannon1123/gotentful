package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
	"net/http"
	"os"
	"time"
	"GotentfulService/common"
	"GotentfulService/assets"
	"GotentfulService/content_types"
	"GotentfulService/entries"
	"GotentfulService/locales"
	"GotentfulService/spaces"
	// "GotentfulService/sync"
	"github.com/julienschmidt/httprouter"
)

//Add Apache Combined Logging
func loggedHandler(h http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, h)
}

//Add Internal Server Error Recovery & Logging
func recoveryHandler(h http.Handler) http.Handler {
	return handlers.RecoveryHandler()(h)
}

func timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 3*time.Minute, "Time Out")
}

func main() {
	//Init Mongo & Other Items
	config := common.StartUp()

	//Initialize Routes
	router := newRouter(config)

	//Chain together Handlers
	chain := alice.New(loggedHandler, recoveryHandler, timeoutHandler).Then(router)

	//Serve it Up
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.Server, config.Port), chain)
}

func newRouter(config *common.Configuration) *httprouter.Router {
	router := httprouter.New()

	assets.RegisterRoutes(router, config.Session, config.Database)
	content_types.RegisterRoutes(router, config.Session, config.Database)
	entries.RegisterRoutes(router, config.Session, config.Database)
	locales.RegisterRoutes(router, config.Session, config.Database)
	spaces.RegisterRoutes(router, config.Session, config.Database)
	//sync.RegisterRoutes(router, config.Session, config.Database)

	return router
}
