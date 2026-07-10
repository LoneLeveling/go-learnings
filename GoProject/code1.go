package main

//Creating an API server in go
//See the http docs: go doc http |grep Listen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product //declaring a slice
func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: homepage")
	fmt.Fprint(w, "Hi Welcome to our first GO HTTP homepage :)")

}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products) //Passed our slice "Products" here
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL.Path)
	vars := mux.Vars(r)
	key := vars["id"]
	// key := r.URL.Path[len("/product/"):] //slicing to fetch just the id field
	for _, Product := range Products {
		if string(Product.Id) == key {
			json.NewEncoder(w).Encode(Product)
		}
	}
}

// Routing
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// http.HandleFunc("/products", returnAllProducts)
	// http.HandleFunc("/", homepage)
	//Get products by Id : Get /products/id
	// http.HandleFunc("/product/", getProductById)
	// http.ListenAndServe("localhost:2000", nil)

	myRouter.HandleFunc("/products", returnAllProducts)
	//Get products by Id : Get /products/id
	myRouter.HandleFunc("/product/{id}", getProductById)
	myRouter.HandleFunc("/", homepage)

	http.ListenAndServe("localhost:2000", myRouter)

}

func main() {
	Products = []Product{
		Product{Id: "1", Name: "Bottles", Quantity: 10, Price: 1000.00},
		Product{Id: "2", Name: "Goggles", Quantity: 4, Price: 4500.00},
	}
	handleRequests()

}
