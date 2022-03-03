package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productJSON := `[
		{
		  "productId": 1,
		  "manufacturer": "Johns-Jenkins",
		  "sku": "p5z343vdS",
		  "upc": "939581000000",
		  "pricePerUnit": "497.45",
		  "quantityOnHand": 9703,
		  "productName": "sticky note"
		},
		{
		  "productId": 2,
		  "manufacturer": "Hessel, Schimmel and Feeney",
		  "sku": "i7v300kmx",
		  "upc": "740979000000",
		  "pricePerUnit": "282.29",
		  "quantityOnHand": 9217,
		  "productName": "leg warmers"
		},
		{
		  "productId": 3,
		  "manufacturer": "Swaniawski, Bartoletti and Bruen",
		  "sku": "q0L657ys7",
		  "upc": "111730000000",
		  "pricePerUnit": "436.26",
		  "quantityOnHand": 5905,
		  "productName": "lamp shade"
		},
		{
		  "productId": 4,
		  "manufacturer": "Runolfsdottir, Littel and Dicki",
		  "sku": "x78426lq1",
		  "upc": "93986215015",
		  "pricePerUnit": "537.90",
		  "quantityOnHand": 2642,
		  "productName": "flowers"
		},
		{
		  "productId": 5,
		  "manufacturer": "Kuhn, Cronin and Spencer",
		  "sku": "r4X793mdR",
		  "upc": "260149000000",
		  "pricePerUnit": "112.10",
		  "quantityOnHand": 6144,
		  "productName": "clamp"
		}
	]`
	err := json.Unmarshal([]byte(productJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	println(r.Method)
	switch r.Method {
	case http.MethodGet:
		productJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJson)
	}
}

func main() {
	fmt.Println("Starting Web Service")
	port := ":5001"
	http.HandleFunc("/product", productHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
