package main

import (
	"fmt"
	"kanggo/absenService/server/v1"
	"log"
	"os"
)
/*
	Created 24 August 2020, by MPutong
	SMS - Service
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
