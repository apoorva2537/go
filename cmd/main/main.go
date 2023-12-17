package main

import (
	"log"
	"net/http"

	"github.com/apoorva2537/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)
func main(){
	r:=mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListAndServe('localhost:8000',r))
}