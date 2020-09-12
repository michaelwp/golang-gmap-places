package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/michaelwp/golang-gmap-places/middlewares"
	"github.com/michaelwp/golang-gmap-places/routers"
	"net/http"
	"time"
)

type Host string

func (h Host) Server() (*http.Server, *mux.Router, string)  {
	// set mux router
	router := mux.NewRouter()
	httpLog := middlewares.RouterLogger(router)

	// set http server
	srv := &http.Server{
		Handler: httpLog,
		Addr: string(h),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	//connecting to router
	routers.Router(router)

	// return server router and response
	resp := fmt.Sprintf("Server running on host %s", string(h))
	return srv, router, resp
}
