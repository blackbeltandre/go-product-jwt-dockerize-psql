package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-product/app"
	"go-product/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/product/new", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/all/product", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/api/product/detail", controllers.GetProductDetail).Methods("GET")
	router.HandleFunc("/api/product/delete", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/product/update", controllers.UpdateProduct).Methods("PATCH")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("port")
	if port == "" {
		port = "8002" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
