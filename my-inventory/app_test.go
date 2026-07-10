package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) { /*NOTE: This testing.M struct has a single method defined in it called Run(),
	which Runs all the other tests present within the package*/

	err := a.Initialise(DBUser, DbPassword, "test") //DbName= "test"
	if err != nil {
		log.Fatal("Error occured while intitialising the database")
	}
	createTable()
	m.Run()
}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS products (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		quantity int,
		price float(10,7),
		PRIMARY KEY (id)
	);`
	_, err := a.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	// Delete all records from the products table.
	_, err := a.DB.Exec("DELETE FROM products")
	if err != nil {
		log.Printf("Error clearing table: %v", err)
	}
	// Reset the auto-increment counter.
	_, err = a.DB.Exec("ALTER TABLE products AUTO_INCREMENT=1")
	if err != nil {
		log.Printf("Error resetting auto_increment: %v", err)
	}
	log.Println("clearTable completed")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT INTO products(name, quantity, price) VALUES('%v', %v, %v)", name, quantity, price)
	_, err := a.DB.Exec(query)
	if err != nil {
		log.Printf("Error adding product: %v", err)
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	// Add a new product. With the auto-increment reset, the new product’s ID will be 1.
	addProduct("keyboard", 200, 500)

	// Create a new GET request for the product with ID 1.
	request, err := http.NewRequest("GET", "/product/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	// Send the request and get the response.
	response := sendRequest(request)

	// Verify the HTTP status code.
	checkStatusCode(t, http.StatusOK, response.Code)
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received: %v", expectedStatusCode, actualStatusCode)
	}
}
func TestCreateProduct(t *testing.T) {
	clearTable()
	var product = []byte(`{"name":"chair", "quantity":1, "price":100}`)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	req.Header.Set("Content-Type", "application/json")

	response := sendRequest(req)
	checkStatusCode(t, http.StatusCreated, response.Code)

	// var m map[string]interface{}
	// json.Unmarshal(response.Body.Bytes(), &m)

	// if m["name"] != "chair" {
	// 	t.Errorf("Expected name: %v, Got: %v", "chair", m["name"])
	// }
	// if m["quantity"] != 1.0 { // Numbers are unmarshaled as float64.
	// 	t.Errorf("Expected quantity: %v, Got: %v", 1.0, m["quantity"])
	// }
}
