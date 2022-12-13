package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hsmyv/go-warehouse/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main(){
	r := mux.NewRouter()

	routes.RegisterWarehouseRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}