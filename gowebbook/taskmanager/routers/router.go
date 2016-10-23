package routers

import(
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router{
	router := mux.NewRouter().StrictSlash(false)
	//routes for User entity
	router = SetUserRoutes(router)
	//routes for Task entity
	router = SetTaskRoutes(router)
	//routes for TaskNote entity
	router = SetNoteRoutes(router)
	return router
}