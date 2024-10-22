package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

var ip, port string

type Product struct {
	Name     string
	Quantity uint
}

func main() {

	RegisterService()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /purchase/{product_id}", func(w http.ResponseWriter, r *http.Request) {
		// discovery service ip: 172.17.0.3
		client, err := rpc.Dial("tcp", "discovery:5555")
		if err != nil {
			panic(err)
		}
		defer client.Close()

		var addr net.TCPAddr
		if err := client.Call("Discovery.Get", "product", &addr); err != nil {
			log.Fatal(err)
		}

		// check if sercice is running or not...
		// w.Write([]byte(addr.String()))

		// get quantity
		res, err := http.Get(fmt.Sprintf("%s%s%s%s", "http://", addr.String(), "/product/", r.PathValue("product_id")))
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		// serialize
		var product Product
		err = json.NewDecoder(res.Body).Decode(&product)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if product.Quantity <= 0 {
			http.Error(w, "product is out of stock", http.StatusNotFound)
			return
		}

		// update quantity
		// this is request to product service for update quantity...
		req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s%s%s", "http://", addr.String(), "/product/", r.PathValue("product_id")), nil)
		if err != nil {
			log.Println("http.NewRequest error:", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Println("httpClient.Do error:", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("order service is healthy\n"))
	})

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("order service is running on container with ip: %s and port: %s\n", ip, port)
	log.Fatal(server.ListenAndServe())
}
