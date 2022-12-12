package routes

import (
	"github.com/gorilla/mux"
	"github.com/akhil/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/product/", 		controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/" , 		controllers.GetProduct).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.GetproductById).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{productId}", controllers.DeleteProduct).Methods("DELETE")
}