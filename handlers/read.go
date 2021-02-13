package handlers

import (
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

func (p *Products) RetrieveAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for READ request")

	listProducts := data.GetAllProducts()
	p.l.Println("[DEBUG] Retrieved all products")

	err := data.ToJSON(listProducts, w)
	if err != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
	}
}

func (p *Products) RetrieveSingle(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for READ request")

	id := getProductID(r)
	p.l.Println("[DEBUG] Retrieved product ID: ", id)

	product, err := data.GetSpecificProduct(id)
	if err != nil {
		http.Error(w, "Unable to find product with the given ID ...", http.StatusInternalServerError)
	}
	p.l.Println("[DEBUG] Retrieved product data successfully!")
	val := data.ToJSON(product, w)
	if val != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
	}
}
