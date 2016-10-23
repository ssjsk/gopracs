package main

import(
	"log"
	"net/http"
	"github.com/codegangsta/negroni"

	"github.com/ssjsk/gowebbook/taskmanager/common"
	"github.com/ssjsk/gowebbook/taskmanager/routers"
)
func main(){
	//Startup
	common.Startup()

	//Get mux router object
	router := routers.InitRoutes()

	//Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServ()
}
