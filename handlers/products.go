package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prateek041/microservices-with-go/data"
)

type Products struct {
	l *log.Logger
}

// used for creating a new isntance of Product list.
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// used to get all the products from the database
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	productList := data.GetData()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error in marshalling the data", http.StatusInternalServerError)
	}
}

// used to add a new product to the database.
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	newProduct := r.Context().Value(ProductKey{}).(data.Product)

	data.AddData(&newProduct) // Adds the product
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	// Now, create a new Product using the data passed in the http request.

	vars := mux.Vars(r)    // all the parameters in the url
	stringId := vars["id"] // id returned is string
	id, err := strconv.Atoi(stringId)
	if err != nil {
		http.Error(rw, "Unable to convert string to integer", http.StatusInternalServerError)
		return
	}

	// we are getting the product from the context present in the middleware
	prod := r.Context().Value(ProductKey{}).(data.Product)

	err = data.UpdateProduct(id, &prod) // add it in the database
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type ProductKey struct{}

func (p Products) MiddlewareTestFunction(next http.Handler) http.Handler {
	// HandlerFunc converts it into a handler function
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) { // this is the handler we are returning
		prod := data.Product{} // an empty table.

		err := prod.FromJSON(r.Body) // create a new entry of type Product.
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		}

		// creates a new context by taking the parent and attaching the key value pair.
		ctx := context.WithValue(r.Context(), ProductKey{}, prod)

		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
