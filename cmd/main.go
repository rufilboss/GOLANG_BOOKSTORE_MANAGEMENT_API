package main

import (
	"log"
	"net/http"

	// "github.com/akhil.go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"rufilboss/crud_api/GOLANG_BOOKSTORE_MANAGEMENT_API/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
