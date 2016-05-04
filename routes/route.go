package routers

import (
	"taskManagerLogin/config"
	"github.com/gorilla/mux"
	"net/http"
	"taskManagerLogin/handlers"
)

func HandleRequests(context config.Context) {
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.Login(context)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", r)
}
