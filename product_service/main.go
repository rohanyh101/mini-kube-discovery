package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Product struct {
	Name     string
	Quantity uint
}

var ip, port string
var products map[string]Product

func main() {

	RegisterService()

	products = map[string]Product{
		"aaa": {Name: "apple", Quantity: 10},
		"bbb": {Name: "banana", Quantity: 7},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("product service is healthy"))
	})

	mux.HandleFunc("PUT /product/{product_id}", func(w http.ResponseWriter, r *http.Request) {
		product := products[r.PathValue("product_id")]
		if product.Quantity <= 0 {
			http.Error(w, "product out of stock", http.StatusBadRequest)
			return
		}

		product.Quantity--
		products[r.PathValue("product_id")] = product
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /product/{product_id}", func(w http.ResponseWriter, r *http.Request) {
		product := products[r.PathValue("product_id")]
		json, err := json.Marshal(product)
		if err != nil {
			log.Println("error marshalling product: ", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	})

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("product service is running on container with ip: %s and port: %s\n", ip, port)
	log.Fatal(server.ListenAndServe())
}
