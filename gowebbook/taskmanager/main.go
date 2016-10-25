package main

import (
	"github.com/codegangsta/negroni"
	"log"
	"net/http"

	"github.com/ssjsk/gowebbook/taskmanager/common"
	"github.com/ssjsk/gowebbook/taskmanager/routers"
)

func main() {
	//Startup
	common.Startup()

	//Get mux router object
	router := routers.InitRoutes()

	//Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
	log.Println(common.AppConfig.Server)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
