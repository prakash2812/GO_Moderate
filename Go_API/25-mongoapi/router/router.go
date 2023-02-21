package router

import (
	controller "github.com/arjun/modules/25-mongoapi/controler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", controller.MarckAsWatched).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteAll", controller.DeleteAllMovie).Methods("DELETE")

	return r

}
