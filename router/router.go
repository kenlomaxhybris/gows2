package router

import (
	"github.com/gorilla/mux"
	"github.com/kenlomaxhybris/goworkshopII/controllers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/workshops", controllers.ReadAllWorkshops).Methods("GET")
	router.HandleFunc("/workshops", controllers.CreateWorkshop).Methods("POST")
	router.HandleFunc("/workshops/{id}", controllers.UpdateWorkshop).Methods("PUT")
	router.HandleFunc("/workshops/{id}", controllers.ReadWorkshop).Methods("GET")
	router.HandleFunc("/workshops/{id}", controllers.DeleteWorkshop).Methods("DELETE")

	return router
}
