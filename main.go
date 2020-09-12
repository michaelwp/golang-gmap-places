package main

import (
	"fmt"
	"github.com/michaelwp/golang-gmap-places/server/v1"
	"log"
	"os"
)
/*
	Created 12 September 2020, by Michael W. Putong
	Google Map - Service
*/

func main(){
	// set host
	host := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	h := v1.Host(host)

	// call the server
	srv, _, resp := h.Server()

	// print log to the screen
	log.Println(resp)
	log.Fatal(srv.ListenAndServe())
}
